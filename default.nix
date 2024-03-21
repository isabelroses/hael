{
  lib,
  buildGoModule,
}:
buildGoModule {
  pname = "hael";
  version = "0.0.1";

  src = ./.;

  #  sha256-0000000000000000000000000000000000000000000=
  vendorHash = null;

  ldflags = ["-s" "-w"];

  meta = {
    description = "Trans rights are human rights";
    homepage = "https://github.com/isabelroses/hael";
    license = with lib.licenses; [mit];
    maintainers = with lib.maintainers; [isabelroses];
    platforms = lib.platforms.all;
  };
}
