package regxgen

var DEFAULT_CONFIG Config

func init() {
	DEFAULT_CONFIG = Config{
		repetetionMax: 5,
	}
}

type Config struct {
	repetetionMax int
}