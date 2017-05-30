# Tag Validate

The ultimate way to validate your data. All you need is add some info on tags with some simple description!

# Install

    `go get `

# Example:

```go
package main

import (
	"fmt"

	"./tagvalidate"
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

Syntex: `{option name}:"{option value}" [{option name2}:"{option value2}"...]`

## For string field you can use those:

| Option    | Value          | Meaning                       |
| --------- | -------------- | ----------------------------- |
| empty     | false/true     | Allow this field empty or not |
| eq        | any string     | Strictly equal to value       |
| neq       | any string     | Strictly not equal to value   |
| starts    |                |                               |
| ends      |                |                               |
| contains  |                |                               |
| ncontains |                |                               |
| upper     |                |                               |
| lower     |                |                               |
| empty     |                |                               |
| len       |                |                               |
| max_len   |                |                               |
| min_len   |                |                               |
| regx      |                |                               |
| type      | See type table |                               |
| func      |                |                               |
|           |                |                               |

Type is a quick access to some frequency data type, most of then is validated by regx. For type, allow those values. To add extra value (if this type allowed extra value) use `,` after type name like: `type:"date,2006-01-02"`:


| Type name | meaning                   | examples       |      Extra value       |
| --------- | ------------------------- | -------------- | :--------------------: |
| int       |                           | 1, +123, -3, 0 |           -            |
| float     |                           |                |           -            |
| hex       |                           |                |           -            |
| ipv4      |                           |                |           -            |
| ip        | ipv4/ipv6                 |                |           -            |
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
| *json     | json dumped string        |                |                        |
| *domain   | domain of a web site      |                |                        |
| *map      |                           |                |                        |
| *list     |                           |                |                        |

\* Those are not done yet. 

## For integers:


| Option | Value      | Meaning                     |
| ------ | ---------- | --------------------------- |
| zero   | false/true | Allow zero or not           |
| eq     | any int    | Strictly equal to value     |
| neq    | any int    | Strictly not equal to value |
| max    |            |                             |
| min    |            |                             |
| func   |            |                             |
|        |            |                             |

## TODO

* Float field.
* Struct field.
* Map field.
* Array field.
* Child struct check.
* Set to default value when empty.