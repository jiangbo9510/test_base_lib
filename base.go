package test_base_lib

var (
	Common = ""
)

func GetCommon() string {
	return Common
}

func SetCommon(v string, x string) {
	Common = v + x
}
