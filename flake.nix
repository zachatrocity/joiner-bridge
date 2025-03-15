{
  description = "Joiner Bridge - A Matrix bridge written in Go";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true;
        };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go environment
            go
            gopls
            go-tools
            golangci-lint
            delve  # Go debugger

            # Database
            postgresql

            # Development tools
            git
            chromium  # Free alternative to Chrome
          ];

          shellHook = ''
            echo "Joiner Bridge Development Environment"
            echo "Go: $(go version)"
          '';

          # Set GOPATH to allow package installation
          GOPATH = "${pkgs.go}/share/go";
        };
      }
    );
}
