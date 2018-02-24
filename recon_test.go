package recon_test

import (
	"github.com/kainoaseto/recon"
	"reflect"
	"testing"
	"time"
)

type ConfigTestSpec struct {
	StringVar string
	IntVar    int
	BoolVar   bool
	DoubleVar float64
	FloatVar  float32
	Template  string
	Timeout   time.Duration
	StrList   []string       `envconfig:"TEST_LIST"`
	StrMap    map[string]int `envconfig:"TEST_MAP"`
}

var ConfigTests = []struct {
	Location interface{}
	Expected interface{}
}{
	{"TEST_STRINVAR", "https://consulstorage.com/kv/hellothere"},
	{"TEST_INTVAR", 45},
	{"TEST_BOOLVAR", true},
	{"TEST_DOUBLEVAR", 45.000432},
	{"TEST_FLOATVAR", float32(0.3)},
	{"export TEST_TEMPLATE", "{{ some/secret/key/path }}"},
	{"TEST_TIMEOUT", 3 * time.Minute},
	{"TEST_LIST", []string{"rob", "ken", "robert"}},
	{"TEST_MAP", map[string]int{"red": 1, "green": 2, "blue": 3}},
}

func TestLoadConfig(t *testing.T) {
	var configSpec ConfigTestSpec
	err := recon.LoadConfig("test.env", &configSpec, "test")
	if err != nil {
		t.Errorf("Failed to load test.env", err)
	}

	configPtr := reflect.ValueOf(&configSpec)
	config := configPtr.Elem()

	for idx, test := range ConfigTests {
		expected := test.Expected

		got := config.Field(idx)

		if !reflect.DeepEqual(got.Interface(), expected) {
			t.Errorf("expected: \"%v\" got: \"%v\"", expected, got)
		}
	}
}
