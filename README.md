# Go Env

Populates environment variables with values defined in a .env File

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Dependencies](#dependencies)

## Installation

```sh
$ go get -u github.com/adnanbrq/goenv
```

## Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/adnanbrq/goenv"
)

func main() {
	/* > .env File

	   PORT=3000
	   PROD=true
	*/

	// false in case you do not want to override existing env. variables
	goenv.Load(false)
	fmt.Println(os.Getenv("PORT")) // prints 3000
}
```

## Dependencies

- [github.com/stretchr/testify - v1.8.0](https://github.com/stretchr/testify)
Assertions