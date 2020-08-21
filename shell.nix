{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  name = "go-dev-environment";
  buildInputs = with pkgs; [
    go
  ];
  shellHook = ''
    echo "Start Go developing..."
  '';
}

