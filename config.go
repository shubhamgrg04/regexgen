package regxgen

var DEFAULT_CONFIG Config

func init() {
	DEFAULT_CONFIG = Config{
		RepetetionMax: 5,
	}
}

type Config struct {
	RepetetionMax int
	Seed int64
}