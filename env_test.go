package envtostruct

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	os.Setenv("ENV_STRING", "test 123")
	os.Setenv("ENV_BOOL", "true")
	os.Setenv("ENV_INT", "1")
	os.Setenv("ENV_INT32", "2147483647")
	os.Setenv("ENV_UINT32", "4294967295")
	os.Setenv("ENV_INT64", "9223372036854775807")
	os.Setenv("ENV_UINT64", "18446744073709551615")
	os.Setenv("ENV_FLOAT32", "3.2")
	os.Setenv("ENV_FLOAT64", "6.4")
	os.Setenv("ENV_SLICE_STRING", "a,b,c")
	os.Setenv("ENV_SLICE_INT", "1,2,3")
	os.Setenv("ENV_SLICE_INT32", "1,2,3")
	os.Setenv("ENV_SLICE_UINT32", "1,2,3")
	os.Setenv("ENV_SLICE_INT64", "1,2,3")
	os.Setenv("ENV_SLICE_UINT64", "1,2,3")
	os.Setenv("ENV_SLICE_FLOAT32", "1.1,1.2,1.3")
	os.Setenv("ENV_SLICE_FLOAT64", "2.1,2.2,2.3")
	os.Setenv("ENV_EMBED_STRING", "x")
	os.Setenv("ENV_EMBED_INT", "1")
	os.Setenv("ENV_SLICE_EMBED_STRING", "z")
	os.Setenv("ENV_SLICE_EMBED_INT", "3")

	type EnvEmbed struct {
		EnvEmbedString string `env:"ENV_EMBED_STRING"`
		EnvEmbedInt    int    `env:"ENV_EMBED_INT"`
	}
	type EnvStruct struct {
		EnvString       string    `env:"ENV_STRING"`
		EnvBool         bool      `env:"ENV_BOOL"`
		EnvInt          int       `env:"ENV_INT"`
		EnvInt32        int32     `env:"ENV_INT32"`
		EnvUint32       uint32    `env:"ENV_UINT32"`
		EnvInt64        int64     `env:"ENV_INT64"`
		EnvUint64       uint64    `env:"ENV_UINT64"`
		EnvFloat32      float32   `env:"ENV_FLOAT32"`
		EnvFloat64      float64   `env:"ENV_FLOAT64"`
		EnvSliceString  []string  `env:"ENV_SLICE_STRING"`
		EnvSliceInt     []int     `env:"ENV_SLICE_INT"`
		EnvSliceInt32   []int32   `env:"ENV_SLICE_INT32"`
		EnvSliceUint32  []uint32  `env:"ENV_SLICE_UINT32"`
		EnvSliceInt64   []int64   `env:"ENV_SLICE_INT64"`
		EnvSliceUint64  []uint64  `env:"ENV_SLICE_UINT64"`
		EnvSliceFloat32 []float32 `env:"ENV_SLICE_FLOAT32"`
		EnvSliceFloat64 []float64 `env:"ENV_SLICE_FLOAT64"`
		EnvSkip         string    `env:"-"`
		EnvAlsoSkip     string
		EnvEmbed        EnvEmbed
	}

	env := EnvStruct{}

	t.Run("Success_Set", func(t *testing.T) {
		if err := Set(&env); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, env.EnvString, "test 123")
		assert.Equal(t, env.EnvBool, true)
		assert.Equal(t, env.EnvInt, int(1))
		assert.Equal(t, env.EnvInt32, int32(2147483647))
		assert.Equal(t, env.EnvUint32, uint32(4294967295))
		assert.Equal(t, env.EnvInt64, int64(9223372036854775807))
		assert.Equal(t, env.EnvUint64, uint64(18446744073709551615))
		assert.Equal(t, env.EnvFloat32, float32(3.2))
		assert.Equal(t, env.EnvFloat64, float64(6.4))
		assert.Equal(t, env.EnvSliceString, []string{"a", "b", "c"})
		assert.Equal(t, env.EnvSliceInt, []int{1, 2, 3})
		assert.Equal(t, env.EnvSliceInt32, []int32{1, 2, 3})
		assert.Equal(t, env.EnvSliceUint32, []uint32{1, 2, 3})
		assert.Equal(t, env.EnvSliceUint64, []uint64{1, 2, 3})
		assert.Equal(t, env.EnvSliceFloat32, []float32{1.1, 1.2, 1.3})
		assert.Equal(t, env.EnvSliceFloat64, []float64{2.1, 2.2, 2.3})
		assert.Equal(t, env.EnvSkip, "")
		assert.Equal(t, env.EnvAlsoSkip, "")
		assert.Equal(t, env.EnvEmbed, EnvEmbed{
			EnvEmbedString: "x",
			EnvEmbedInt:    1,
		})
	})
}
