package envparse

import (
	"fmt"
	"testing"
)

var testData = []string{
	"APP_INT1=1",
	"APP_INT2=2",
	"APP_UINT1=123",
	"APP_UINT2=456",
	"APP_FLOAT1=456e-3",
	"APP_BOOL1=TRUE",
	"APP_STRINGS_0=String 1",
	"APP_STRINGS_1=String 2",
	"APP_STRUCT1_0_INT1=123456",
	"APP_STRUCT2_INT1=7890",
}

type conf struct {
	Int1    *int          `env:"name=INT1"`
	Int2    int64         `env:"name=INT2"`
	Uint1   uint          `env:"name=UINT1"`
	Uint2   uint64        `env:"name=UINT2"`
	Float1  float64       `env:"name=FLOAT1"`
	Bool1   bool          `env:"name=BOOL1"`
	String1 string        `env:"name=STRING1"`
	Strings []string      `env:"name=STRINGS"`
	Struct1 []*SubStruct1 `env:"name=Struct1"`
	//Struct2 []*SubStruct2 `env:"name=Struct2,required"`
}

type SubStruct1 struct {
	Int1    int     `env:"name=INT1"`
	Uint1   uint    `env:"name=UINT1"`
	Float1  float64 `env:"name=float1"`
	Bool1   bool    `env:"name=bool1"`
	String1 string  `env:"name=str1"`
	//Struct2 *SubStruct2 `env:"name=struct2"`
}

type SubStruct2 struct {
	Int1    int     `env:"name=INT1"`
	Uint1   uint    `env:"name=UINT1"`
	Float1  float64 `env:"name=float1"`
	Bool1   bool    `env:"name=bool1"`
	String1 string  `env:"name=str1"`
	//Struct1 *SubStruct1 `env:"name=struct1"`
}

//var expectedData = &conf{
//	Int1:    1,
//	Int2:    2,
//	Uint1:   123,
//	Uint2:   456,
//	Float1:  456e-3,
//	Bool1:   true,
//	String1: "Hello world I am alive",
//	Struct1: SubStruct1{SubInt1: 123456},
//	Struct2: &SubStruct1{SubInt1: 7890},
//}

func TestParse(t *testing.T) {
	c := &conf{}
	err := Parse(c, testData)
	if err != nil {
		t.Error(err.(*errorList).FullError())
	}
	fmt.Println(c)
}
