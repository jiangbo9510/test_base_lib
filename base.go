package test_base_lib

var (
	Common = ""
)

func GetString() (string, string) {
	return "v", "0.0.2"
}

func GetCommon() string {
	return Common
}

func SetCommon(v string) {
	Common = v
}
