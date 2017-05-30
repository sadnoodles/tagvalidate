# Tag Validate

The ultimate way to validate your data. All you need is add some info on tags with some simple description!

# Install

    `go get github.com/sadnoodles/tagvalidate`

# Example:

```go
package main

import (
	"fmt"

	"github.com/sadnoodles/tagvalidate"
)

type Common struct {
	Id   int64  `zero:"true" max:"20" min:"7" neq:"11"` // Add tags here
	Uuid string `empty:"true" type:"uuid"` // Add tags here
}

func main() {
	c := new(Common)
	c.Id = 4         //Validate error. @Field:Id tag:min wanted:7, Got: 4
	c.Uuid = "ssasd" //Validate error. @Field:Uuid tag:type wanted:uuid, Got: "ssasd"
	err := tagvalidate.Check(c) // Check a instance of Common struct
	if err != nil {
		fmt.Println(err)
	}
}
```

# Usage:

Syntax: `{option name}:"{option value}" [{option name2}:"{option value2}"...]`

## For string field you can use those:

| Option    | Value              | Meaning                            |
| --------- | ------------------ | ---------------------------------- |
| empty     | false/true         | Allow this field empty or not      |
| eq        | any string (s)     | Strictly equal to s                |
| neq       | any string         | Strictly not equal to s            |
| starts    | any string         | Starts with s                      |
| ends      | any string         | Ends with s                        |
| contains  | any string         | Contains s                         |
| ncontains | any string         | Note contains s                    |
| upper     | false/true         |                                    |
| lower     | false/true         |                                    |
| len       | int                | Must be this long                  |
| max_len   | int                | Max length                         |
| min_len   | int                | Min length                         |
| regx      | reg exp            | re check                           |
| **type**  | **See type table** | **Frequency types**                |
| func      | Func name          | Custom functions under your struct |
|           |                    | func (string)  bool                |

Type is a quick access to some frequency data type, most of then is validated by regx.  Use type tag like this:

```go
type LoginLog struct{
    User string `empty:"false" type:"email"`
    IP string `type:"ipv4"` 
}
```


For type, allow those values. To add extra value (if this type allowed extra value) use `,` after type name like: `type:"date,2006-01-02"`:


| Type name | meaning                   | examples       |      Extra value       |
| --------- | ------------------------- | -------------- | :--------------------: |
| int       |                           | 1, +123, -3, 0 |           -            |
| float     |                           |                |           -            |
| hex       |                           |                |           -            |
| ipv4      | ipv4 only                 |                |           -            |
| ip        | ipv4 or ipv6              |                |           -            |
| email     |                           |                |           -            |
| url       |                           |                |           -            |
| hexcolor  | color use hex             |                |           -            |
| fullpath  | windows or unix full path |                |           -            |
| uuid3     |                           |                |           -            |
| uuid4     |                           |                |           -            |
| uuid5     |                           |                |           -            |
| uuid      |                           |                |           -            |
| num       | number                    |                |           -            |
| alpha     | alpha table               |                |           -            |
| md5       | MD5 checksum (length 32)  |                |           -            |
| md5(16)   | MD5 checksum (length 16)  |                |           -            |
| base64    | base64 encoded string     |                |           -            |
| date      | date format               |                | date formating  string |
| json      | json dumped string        |                |                        |
| *domain   | domain of a web site      |                |                        |
| *map      |                           |                |                        |
| *list     |                           |                |                        |

\* Those are not done yet. 

To check if data is a LAN IP:
```go
InnerIP string `type:"ip" regx:"^(127|192|172|10[\\D])"`
```
NOTE: dot(".") in tag will cause tag parse error. Also you should use two `\\` to escape slash.

## For integers:


| Option | Value       | Meaning                            |
| ------ | ----------- | ---------------------------------- |
| zero   | false/true  | Allow zero or not                  |
| eq     | any int (i) | Strictly equal to i                |
| neq    | any int     | Strictly not equal to i            |
| max    | int         | Max value                          |
| min    | int         | Min value                          |
| func   | Func name   | Custom functions under your struct |
|        |             | func (int64)  bool                 |

## TODO

* Float field.
* Struct field.
* Map field.
* Array field.
* Child struct check.
* Set to default value when empty.