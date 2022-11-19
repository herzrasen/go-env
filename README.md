# go-env

Tired of writing the same env helpers over and over again? Yeah. me too. That's
why this repo exists. 

## Installation

```shell
go get github.com/herzrasen/go-env
```

## Usage

To get an error when no environment variable `key` exists:

```go
import "github.com/herzrasen/go-env/env"

v, err := env.GetOrError("FOO")
if err != nil {
	return err
}
fmt.Printf("%s", v)
```

To either load the environment variable or use a default value:

```go
import "github.com/herzrasen/go-env/env"

v := env.GetOrElse("FOO", "baz")
fmt.Printf("%s", v)
```