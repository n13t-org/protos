
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


