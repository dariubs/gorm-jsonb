gorm jsonb
==========

[gorm](https://gorm.io) jsonb datatype for postgres

install
-------

```
go get github.com/dariubs/gorm-jsonb
```

example
-------

```go
package main

import "github.com/dariubs/gorm-jsonb"

type Data struct {
	Name   string
	Social gormjsonb.JSONB
}

func main() {

	data := Data{}

	data.Social = make(map[string]interface{})

	data.Social["github"] = "github.com/dariubs"
	data.Social["instagram"] = "instagram.com/drawush"

}

```

license
-------

GNU GPL v3