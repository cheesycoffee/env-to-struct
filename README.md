# env-to-struct
OS Environment Values To Struct using "env" tag

> **supported data types :**
  - string
  - bool
  - int
  - int32
  - int64
  - uint32
  - uint64
  - float32
  - float64
  - embeded struct with env tag
  - []string
  - []int
  - []int32
  - []uint32
  - []int64
  - []uint32
  - []float32
  - []float64
  - map[string]interface{}

> **example Usage :**
```go
import (
	"log"

	envtostruct "github.com/cheesycoffee/env-to-struct"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
    
    type EnvEmbed struct {
        Data string `env:"ENV_EMBED_DATA"`
    }
    
    type Env struct{
		EnvString string `env:"ENV_STRING"`
        EnvSliceString []string `env:"ENV_SLICE_STRING"`
        EnvSliceInt []int `env:"ENV_SLICE_INT"`
        EnvMapString map[string]interface{} `env:"ENV_MAP_STRING"`
        EnvEmbed EnvEmbed
	}

	env := Env{}
	if err := envtostruct.Set(&env); err != nil {
		log.Fatal(err)
	}
    
}
```

> **example .env file :**
```
ENV_STRING="abc"
ENV_SLICE_STRING="a,b,c"
ENV_SLICE_INT="1,2,3"
ENV_MAP_STRING='{"id" : 123, "name" : "abc", "isActive" : true, "account" : [{"id" : 321, "amount" : 123.45}]}'
ENV_EMBED_DATA="xyz"
```

> **see unit test for all supported types**