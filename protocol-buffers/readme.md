Scalar Types
=

```protobuf
// Numbers, 
// default: 0 
int32, int64, sint32, sint64 
uint32, uint64
fixed32, fixed64, sfixed32, sfixed64
float, double

// Boolean, 
// default: false
bool 

// String, 
// default: empty string 
// - UTF-8 or 7 bit ASCII
string

// Bytes, 
// default: empty byte
// - images etc.
bytes

// Repeated, 
// default: empty list 
// - 0 or more value
repeated <type> <name> = <tag>

// Enum, 
// default: first value defined 
// - first tag must be 0 not 1

__________________________________________

// Example proto
syntax = "proto3";

enum EyeColor {
  EYE_COLOR_UNKNOWN = 0;
  EYE_COLOR_BROWN = 1;
  EYE_COLOR_GREEN = 2;
  EYE_COLOR_BLUE = 3;
}

message Account{
  uint32 id = 1;
  string name = 2;
  bytes image = 3;
  bool is_verified = 4;
  float height = 5;
  
  repeated string phones = 6;
  EyeColor eye_color = 7;
}
```

Tags
=

- Field names are not important for serialization but tags are.
- Between 1 - 536870911
- Google reserved some tags, 19000 - 19999
- 1-15 -> 1 byte, 16-2047 -> 2 bytes, ...

- If a field removed, do not use its tag.
- If you need to rename a field, just rename it and don't change the tag.
- If a field cannot interchangeable while you want to change a field's type, check the documentation or remove the field and don't use its tag and add another field as desired with a new tag.

```protobuf
message Test {
    reserved 2,15,9 to 11; // 2,15,9,10,11 tags reserved and cannot be assigned.
    reserved "first_name", "last_name"; // first_name and last_name names reserved and cannot be assigned.
    uint32 id = 1;
}
```

Tips
=
```protobuf
// No negative value
uint64
uint32

// Accept negative value 
// (but not efficient serializing them)
// (10 bytes long)
int32
int64

// Accept negative value 
// (less efficient serializing positive values)
sint32
sint64

// Always 4 bytes long
fixed32
sfixed32

// Always 8 bytes long
fixed64
sfixed64
```

<h4>Oneof</h4>
- Cannot be repeated.
- Evolving the schema is complicated.

<h4>Maps</h4>
- Cannot be repeated.
- Cannot use float, double and enums as key.
- No ordering.

<h4>Timestamp and Duration</h4>

```protobuf
syntax = "proto3";

import "google/protobuf/duration.proto";
import "google/protobuf/timestamp.proto";

message Account {
  google.protobuf.Timestamp created_at = 1;
  google.protobuf.Duration validity = 2;
}
```

Good to have
=
- Name of the file should be lowered snake case (test_name).
- Each line 80 chars max.
- Indentation 2 spaces.
- Double quotes for strings. ("test-string")
- License header above the syntax, simple description after, syntax, package, imports and options.
- Enums, messages and services if any.
- Imports should be alphabetically.
- Service, enum, message and rpc names should be pascal case (TestName).
- Enum fields should be upper snake case (TEST_NAME). And zero-value enum: XXX_UNSPECIFIED
- Repeated fields should be plural.
- Message field names lower snak case (test_field_name).
- Package names lowercase and no whitespace (test.package).
- THERE IS ALSO UBER PROTOBUF STYLE

Services
=

```protobuf
// Not designed for serialization and deserialization.
// It is designed for communication.
// Main framework for communication is gRPC.
// Set of endpoints, that are define an API.


message GetSomethingRequest{}
message GetSomethingResponse{}
message ListSomethingRequest{}
message ListSomethingResponse{}

service MyService{
  rpc GetSomething(GetSomethingRequest) returns (GetSomethingResponse);
  rpc ListSomething(ListSomethingRequest) returns (ListSomethingResponse);
}
```






























