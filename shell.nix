#with import <nixpkgs> { };
with import ~/src/nixpkgs { };

runCommand "dummy" {
  buildInputs = [
    go_1_8
    stdenv
    stdenv.cc
    protobuf
    bazel
  ];
  shellHook = ''
    unset SSL_CERT_FILE
    export GOPATH=$(readlink -f ../../../..)
    export PATH=$GOPATH/bin:$PATH
    PATH=$(readlink -f ../../bin):$PATH
  '';
} ""
