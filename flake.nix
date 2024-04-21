{
  description = "raGBAg";

  inputs = {
    nixpkgs = {
      type = "github";
      owner = "NixOS";
      repo = "nixpkgs";
      ref = "nixos-unstable";
    };

    futils = {
      type = "github";
      owner = "numtide";
      repo = "flake-utils";
      ref = "main";
    };

    ryuki = {
      type = "github";
      owner = "iRyukizo";
      repo = "nixos-config";
      ref = "master";
    };
  };

  outputs =
    { self
    , futils
    , nixpkgs
    , ryuki
    }@inputs:
    let
      inherit (futils.lib) eachDefaultSystem;

      version = builtins.substring 0 8 (self.lastModifiedDate or self.lastModified or "19700101");
    in
    eachDefaultSystem (system:
    let
      pkgs = import nixpkgs { inherit system; };
      ryukipkgs = ryuki.packages.${system};
    in
    rec {
      packages = rec {
        ragbag = pkgs.buildGoModule {
          pname = "ragbag";
          inherit version;
          src = ./.;
          vendorHash = null;
        };
        default = ragbag;
      };

      devShells.default = pkgs.mkShell {
        name = "ragbag dev shell";
        nativeBuildInputs = with pkgs; [
          go
          ryukipkgs.gopkgsite
        ];
      };
    }
    );
}
