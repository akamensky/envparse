package envparse

import (
	"os"
	"reflect"
	"testing"
)

func TestParse_1(t *testing.T) {
	oldPrefix := defaultPrefix
	SetPrefix("")
	defer func() {
		defaultPrefix = oldPrefix
	}()

	type testType struct {
		p1 string `env:"name=hello"`
	}
	c := testType{}

	err := Parse(nil, nil)
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
	err = Parse(nil, []string{})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
	err = Parse(nil, []string{"BAD"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
	err = Parse(c, nil)
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
	err = Parse(nil, []string{"TEST=1"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
	err = Parse(c, []string{"TEST=1"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
	err = Parse(&c, []string{"TEST=1"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}

	type testType2 struct {
		P1 string `env:"name=hello"`
	}
	c2 := testType2{}

	err = Parse(&c2, []string{"TEST=1"})
	if err != nil {
		t.Errorf("expected err == nil, but got '%s'", FullError(err))
	}
}

func TestParse_2(t *testing.T) {
	type testType struct {
		P1 string `env:"name=hello"`
		p2 string `env:"name=test0"`
		P2 string `env:"name=test1,default=12,required"`
	}
	c2 := &testType{}

	err := Parse(c2, []string{"APP_TEST=1"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
}

func TestParse_3(t *testing.T) {
	type testType struct {
		P0 *testType `env:"name=hello0"`
	}
	c := &testType{}
	err := Parse(c, []string{"APP_HELLO0=1"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}

	type testType1 struct {
		P1 *testType `env:"name=hello1,required"`
		P2 string    `env:"name=hello2,required"`
	}
	c1 := &testType1{}
	err = Parse(c1, []string{"APP_HELLO3=1"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}

	type testType2 struct {
		P3 complex64 `env:"name=hello3"`
	}
	c2 := &testType2{}
	err = Parse(c2, []string{"APP_HELLO3=1"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
}

func TestParse_4(t *testing.T) {
	type testType struct {
		P0 []string  `env:"name=hello0"`
		P1 *testType `env:"name=embed"`
	}
	c := &testType{}
	err := Parse(c, []string{"APP_HELLO0_00=1"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}

	type testType1 struct {
		P1 []*testType1 `env:"name=hello1"`
		P2 string       `env:"name=hello2"`
	}
	c1 := &testType1{}
	err = Parse(c1, []string{"APP_HELLO1_00=1"})
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}

	type testType2 struct {
		P3 []complex64 `env:"name=hello3"`
	}
	c2 := &testType2{}
	err = Parse(c2, []string{"APP_HELLO3_00=1"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
}

func TestParse_5(t *testing.T) {
	type testType struct {
		Int   int   `env:"name=int"`
		Int8  int8  `env:"name=int8"`
		Int16 int16 `env:"name=int16"`
		Int32 int32 `env:"name=int32"`
		Int64 int64 `env:"name=int64"`
	}
	c := &testType{}
	e := &testType{
		Int:   0,
		Int8:  0,
		Int16: 0,
		Int32: 0,
		Int64: 0,
	}
	err := Parse(c, nil)
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}
	if !reflect.DeepEqual(e, c) {
		t.Errorf("expected '%v', but got '%v'", e, c)
	}

	type testType1 struct {
		Int   int   `env:"name=int,default=1"`
		Int8  int8  `env:"name=int8,default=2"`
		Int16 int16 `env:"name=int16,default=3"`
		Int32 int32 `env:"name=int32,default=4"`
		Int64 int64 `env:"name=int64,default=5"`
	}
	c1 := &testType1{}
	e1 := &testType1{
		Int:   1,
		Int8:  2,
		Int16: 3,
		Int32: 4,
		Int64: 5,
	}
	err = Parse(c1, nil)
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}
	if !reflect.DeepEqual(e1, c1) {
		t.Errorf("expected '%v', but got '%v'", e1, c1)
	}

	type testType2 struct {
		Int   int   `env:"name=int,default=1"`
		Int8  int8  `env:"name=int8,default=2"`
		Int16 int16 `env:"name=int16,default=3"`
		Int32 int32 `env:"name=int32,default=4"`
		Int64 int64 `env:"name=int64,default=5"`
	}
	c2 := &testType2{}
	e2 := &testType2{
		Int:   0,
		Int8:  0,
		Int16: 0,
		Int32: 0,
		Int64: 0,
	}
	err = Parse(c2, []string{"APP_INT=a", "APP_INT8=a", "APP_INT16=a", "APP_INT32=a", "APP_INT64=a"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
	if !reflect.DeepEqual(e2, c2) {
		t.Errorf("expected '%v', but got '%v'", e2, c2)
	}
}

func TestParse_6(t *testing.T) {
	type testType struct {
		Uint   uint   `env:"name=uint"`
		Uint8  uint8  `env:"name=uint8"`
		Uint16 uint16 `env:"name=uint16"`
		Uint32 uint32 `env:"name=uint32"`
		Uint64 uint64 `env:"name=uint64"`
		Byte   byte   `env:"name=byte"`
	}
	c := &testType{}
	e := &testType{
		Uint:   0,
		Uint8:  0,
		Uint16: 0,
		Uint32: 0,
		Uint64: 0,
		Byte:   0,
	}
	err := Parse(c, nil)
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}
	if !reflect.DeepEqual(e, c) {
		t.Errorf("expected '%v', but got '%v'", e, c)
	}

	type testType1 struct {
		Uint   uint   `env:"name=uint,default=1"`
		Uint8  uint8  `env:"name=uint8,default=2"`
		Uint16 uint16 `env:"name=uint16,default=3"`
		Uint32 uint32 `env:"name=uint32,default=4"`
		Uint64 uint64 `env:"name=uint64,default=5"`
		Byte   byte   `env:"name=byte,default=255"`
	}
	c1 := &testType1{}
	e1 := &testType1{
		Uint:   1,
		Uint8:  2,
		Uint16: 3,
		Uint32: 4,
		Uint64: 5,
		Byte:   255,
	}
	err = Parse(c1, nil)
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}
	if !reflect.DeepEqual(e1, c1) {
		t.Errorf("expected '%v', but got '%v'", e1, c1)
	}

	type testType2 struct {
		Uint   uint   `env:"name=uint,default=1"`
		Uint8  uint8  `env:"name=uint8,default=2"`
		Uint16 uint16 `env:"name=uint16,default=2"`
		Uint32 uint32 `env:"name=uint32,default=3"`
		Uint64 uint64 `env:"name=uint64,default=4"`
		Byte   byte   `env:"name=byte,default=255"`
	}
	c2 := &testType2{}
	e2 := &testType2{
		Uint:   0,
		Uint8:  0,
		Uint16: 0,
		Uint32: 0,
		Uint64: 0,
		Byte:   0,
	}
	err = Parse(c2, []string{"APP_UINT=a", "APP_UINT8=a", "APP_UINT16=a", "APP_UINT32=a", "APP_UINT64=a", "APP_BYTE=a"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
	if !reflect.DeepEqual(e2, c2) {
		t.Errorf("expected '%v', but got '%v'", e2, c2)
	}
}

func TestParse_7(t *testing.T) {
	type testType struct {
		Float32 float32 `env:"name=float32"`
		Float64 float64 `env:"name=float64"`
	}
	c := &testType{}
	e := &testType{
		Float32: 0,
		Float64: 0,
	}
	err := Parse(c, nil)
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}
	if !reflect.DeepEqual(e, c) {
		t.Errorf("expected '%v', but got '%v'", e, c)
	}

	type testType1 struct {
		Float32 float32 `env:"name=float32,default=1.2"`
		Float64 float64 `env:"name=float64,default=10e2"`
	}
	c1 := &testType1{}
	e1 := &testType1{
		Float32: 1.2,
		Float64: 10e2,
	}
	err = Parse(c1, nil)
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}
	if !reflect.DeepEqual(e1, c1) {
		t.Errorf("expected '%v', but got '%v'", e1, c1)
	}

	type testType2 struct {
		Float32 float32 `env:"name=float32,default=1.2"`
		Float64 float64 `env:"name=float64,default=10e2"`
	}
	c2 := &testType2{}
	e2 := &testType2{
		Float32: 0,
		Float64: 0,
	}
	err = Parse(c2, []string{"APP_FLOAT32=a", "APP_FLOAT64=a", "APP_UINT16=a", "APP_UINT32=a", "APP_UINT64=a", "APP_BYTE=a"})
	if err == nil {
		t.Errorf("expected err != nil, but got err == nil")
	}
	if !reflect.DeepEqual(e2, c2) {
		t.Errorf("expected '%v', but got '%v'", e2, c2)
	}
}

func TestParse_8(t *testing.T) {
	type testType struct {
		Bool bool `env:"name=bool"`
	}
	c := &testType{}
	e := &testType{
		Bool: false,
	}
	err := Parse(c, nil)
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}
	if !reflect.DeepEqual(e, c) {
		t.Errorf("expected '%v', but got '%v'", e, c)
	}

	type testType1 struct {
		Bool bool `env:"name=bool,default=111"`
	}
	c1 := &testType1{}
	e1 := &testType1{
		Bool: true,
	}
	err = Parse(c1, nil)
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}
	if !reflect.DeepEqual(e1, c1) {
		t.Errorf("expected '%v', but got '%v'", e1, c1)
	}

	type testType2 struct {
		Bool bool `env:"name=bool"`
	}
	c2 := &testType2{}
	e2 := &testType2{
		Bool: false,
	}
	err = Parse(c2, []string{"APP_BOOL=f", "APP_FLOAT64=a", "APP_UINT16=a", "APP_UINT32=a", "APP_UINT64=a", "APP_BYTE=a"})
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}
	if !reflect.DeepEqual(e2, c2) {
		t.Errorf("expected '%v', but got '%v'", e2, c2)
	}
}

func TestParse_9(t *testing.T) {
	_ = os.Setenv("APP_TEST", "1")
	_ = os.Setenv("APP_TEST", "1")
	type testType struct {
		I int `env:"name=TEST,required"`
	}
	c := &testType{}
	origUnsetEnv := unsetEnv
	unsetEnv = true
	err := Parse(c, os.Environ())
	unsetEnv = origUnsetEnv
	if err != nil {
		t.Errorf("expected err == nil, but got '%v'", FullError(err))
	}
	if c.I != 1 {
		t.Errorf("expected c.I == 1, but got c.I == %d", c.I)
	}
	if os.Getenv("APP_TEST") != "" {
		t.Errorf("expected APP_TEST == '', but got APP_TEST == '%s'", os.Getenv("APP_TEST"))
	}
}

func TestRequiredSliceOfStrings(t *testing.T) {
	env := []string{
		"APP_STRINGS_0=1",
	}
	type conf struct {
		Strings []string `env:"name=strings,required"`
	}
	c := &conf{}
	err := Parse(c, env)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(c.Strings) != 1 || c.Strings[0] != "1" {
		t.Errorf("expected c.Strings[0] == '1', but got: %v", c.Strings)
	}
}

func TestRequiredSliceOfInts(t *testing.T) {
	env := []string{
		"APP_INTS_0=1",
	}
	type conf struct {
		Ints []int `env:"required"`
	}
	c := &conf{}
	err := Parse(c, env)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(c.Ints) != 1 || c.Ints[0] != 1 {
		t.Errorf("expected c.Ints[0] == 1, but got: %v", c.Ints)
	}
}
