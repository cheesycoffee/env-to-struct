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

	type Env struct{
		EnvString string `env:"ENV_STRING"`
	}

	env := Env{}
	if err := envtostruct.Set(&env); err != nil {
		log.Fatal(err)
	}

    fmt.Println(env.EnvString)
}
```

> **see unit test for all supported types**

**for slices use string comma separated. for example "1,2,3," for []int**