{
  description = "Joiner Manager Bridge Development Environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        python = pkgs.python3;
        pythonPackages = python.pkgs;
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Python environment
            (python.withPackages (ps: with ps; [
              # Core dependencies
              ps.mautrix
              ps.psycopg2  # For PostgreSQL support
              
              # Development tools
              ps.black
              ps.pylint
              ps.mypy
              ps.pytest
            ]))

            # Database
            postgresql

            # Development tools
            git
          ];

          shellHook = ''
            echo "Joiner Manager Bridge Development Environment"
            echo "Python: ${python.version}"
          '';
        };
      }
    );
}
