{
  description = "A Nix flake with Go and TypeScript";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go dependencies
            go
            gopls  # Go language server for IDE support
            go-tools  # Staticcheck, etc.

            # TypeScript/Node.js dependencies
            nodejs_20  # Node.js (version 20 as an example)
            nodePackages.typescript  # TypeScript compiler
            nodePackages.ts-node  # For running TS directly
            nodePackages.pnpm  # Optional: a fast package manager
          ];

          shellHook = ''
            echo "Welcome to the go2ts dev environment!"
            echo "Go version: $(go version)"
            echo "Node version: $(node --version)"
            echo "TypeScript version: $(tsc --version)"
          '';
        };
      });
}
