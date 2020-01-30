package releasetest

type ReleaseTest struct {
	name string
	age  int
}
type Config struct {
	name string
	age  int
}
type Option func(*Config)

func OptionName(name string) func(*Config) {
	return func(config *Config) {
		config.name = name
	}
}
func OptionAge(age int) func(*Config) {
	return func(config *Config) {
		config.age = 100
	}
}

func New(options ...Option) *ReleaseTest {
	c := new(Config)

	for _, opt := range options {
		opt(c)
	}

	return &ReleaseTest{
		name: c.name,
		age:  c.age,
	}
}
