{ pkgs, lib, config, inputs, ... }:

{

  packages = with pkgs; [
    git
  ];

  languages.go.enable = true;

}
