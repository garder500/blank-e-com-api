package api

import (
	"ecom/internal/models"
	"ecom/internal/utils"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthRequestFormat struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseFormat struct {
	Token string `json:"token"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthRequestFormat
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.NotHandledParseError(w, r, err)
		return
	}

	if req.Username == "" || req.Password == "" {
		utils.ReplyError(w, http.StatusBadRequest, utils.FormattedResponse{
			Message: "Invalid request",
			Code:    http.StatusBadRequest,
			Details: []utils.DetailsResponse{
				{
					Code:    http.StatusBadRequest,
					Message: "Username and password are required",
				},
			},
		})
		return
	}

	// Retrieve the database from the context
	db, err := utils.GetDBFromReq(r)
	if err != nil {
		utils.ReplyError(w, http.StatusInternalServerError, utils.FormattedResponse{
			Message: "Database connection error",
			Code:    http.StatusInternalServerError,
			Details: []utils.DetailsResponse{
				{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			},
		})
		return
	}

	// let's search if a User with this Username and Password exists
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := db.Where(user).First(user).Error; err != nil {
		utils.ReplyError(w, http.StatusUnauthorized, utils.FormattedResponse{
			Message: "Authentication failed",
			Code:    http.StatusUnauthorized,
			Details: []utils.DetailsResponse{
				{
					Code:    http.StatusUnauthorized,
					Message: "Invalid username or password",
				},
			},
		})
		return
	}

	// Let's see if the username is already taken
	userExists := &models.User{}
	userExistsErr := db.Where("username = ?", req.Username).First(userExists).Error
	if userExistsErr == nil {
		utils.ReplyError(w, http.StatusConflict, utils.FormattedResponse{
			Message: "Database error",
			Code:    http.StatusConflict,
			Details: []utils.DetailsResponse{
				{
					Code:    http.StatusConflict,
					Message: "User already exists",
				},
			},
		})
		return
	}else if userExistsErr != gorm.ErrRecordNotFound {
		utils.ReplyError(w, http.StatusInternalServerError, utils.FormattedResponse{
			Message: "Database error",
			Code:    http.StatusInternalServerError,
			Details: []utils.DetailsResponse{
				{
					Code:    http.StatusInternalServerError,
					Message: userExistsErr.Error(),
				},
			},
		})
		return
	}

	// let's insert the user into the database
	if err := db.Create(user).Error; err != nil {
		utils.ReplyError(w, http.StatusInternalServerError, utils.FormattedResponse{
			Message: "Database error",
			Code:    http.StatusInternalServerError,
			Details: []utils.DetailsResponse{
				{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			},
		})
		return
	}

	token, err := utils.GenerateJWTToken(&utils.JwtCustomClaims{
		Claims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(user.ID), 10),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "ecom",
			Audience:  []string{"ecom_users"},
		},
		Role: user.Roles,
	}, os.Getenv("JWT_SECRET_KEY"))

	if err != nil {
		utils.ReplyError(w, http.StatusInternalServerError, utils.FormattedResponse{
			Message: "Token generation error",
			Code:    http.StatusInternalServerError,
			Details: []utils.DetailsResponse{
				{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			},
		})
		return
	}

	utils.ReplyJson(w, http.StatusOK, AuthResponseFormat{
		Token: token,
	})
}
