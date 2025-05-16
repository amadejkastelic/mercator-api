{ pkgs, version, ... }:
pkgs.buildGoModule {
  pname = "mercator-api-cli";
  version = version;
  src = ./..;
  vendorHash = null;
  doCheck = true;
  subPackages = [ "cmd/mercator-api-cli" ];
}
