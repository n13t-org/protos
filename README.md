
# Usage
```go
import (
    ...
    gitlab "gitlab.com/n13t/protos/pkg/n13t/gitlab/v4"
    ...
)
```

- Auto generated stubs only support for Golang. 
> If you use other languages you may want to clone this repo and gen stubs yourself or fork then build with CI, etc.



# [protoc-gen-validate (PGV)](https://github.com/envoyproxy/protoc-gen-validate)
```
syntax = "proto3";

package examplepb;

import "validate/validate.proto";

message Person {
  uint64 id    = 1 [(validate.rules).uint64.gt    = 999];

  string email = 2 [(validate.rules).string.email = true];

  string name  = 3 [(validate.rules).string = {
                      pattern:   "^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$",
                      max_bytes: 256,
                   }];

  Location home = 4 [(validate.rules).message.required = true];

  message Location {
    double lat = 1 [(validate.rules).double = { gte: -90,  lte: 90 }];
    double lng = 2 [(validate.rules).double = { gte: -180, lte: 180 }];
  }
}
```

in/not_in: these two rules permit specifying white/blacklists for the values of a field.
```
// x must be either "foo", "bar", or "baz"
string x = 1 [(validate.rules).string = {in: ["foo", "bar", "baz"]}];

// x cannot be "fizz" nor "buzz"
string x = 1 [(validate.rules).string = {not_in: ["fizz", "buzz"]}];
```

well-known formats: these rules provide advanced constraints for common string patterns. These constraints will typically be more permissive and performant than equivalent regular expression patterns, while providing more explanatory failure descriptions.
```
// x must be a valid email address (via RFC 1034)
string x = 1 [(validate.rules).string.email = true];

// x must be a valid address (IP or Hostname).
string x = 1 [(validate.rules).string.address = true];

// x must be a valid hostname (via RFC 1034)
string x = 1 [(validate.rules).string.hostname = true];

// x must be a valid IP address (either v4 or v6)
string x = 1 [(validate.rules).string.ip = true];

// x must be a valid IPv4 address
// eg: "192.168.0.1"
string x = 1 [(validate.rules).string.ipv4 = true];

// x must be a valid IPv6 address
// eg: "fe80::3"
string x = 1 [(validate.rules).string.ipv6 = true];

// x must be a valid absolute URI (via RFC 3986)
string x = 1 [(validate.rules).string.uri = true];

// x must be a valid URI reference (either absolute or relative)
string x = 1 [(validate.rules).string.uri_ref = true];

// x must be a valid UUID (via RFC 4122)
string x = 1 [(validate.rules).string.uuid = true];
```
