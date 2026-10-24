# Changelog

## 1.0.1 - 2026-10-24

- Build with Go 1.27.3, fixing vulnerabilities in the standard library,
  including certificate verification bypasses and denial of service issues
  in `crypto/x509` and `crypto/tls`, and HTTP/2 memory exhaustion in
  `net/http`
- Update `google.golang.org/grpc` to v1.83.2, fixing HTTP/2 memory
  exhaustion and other HTTP/2 transport issues
- Update `golang.org/x/net` to v0.60.0, fixing a crash in the HTTP/2 HPACK
  encoder

## 1.0.0 - 2026-04-04

_Initial release_
