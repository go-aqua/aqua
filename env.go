package nova // import "github.com/novakit/nova"

// Env nil-safe environment string
type Env string

const (
	// Production production env value
	Production = Env("production")

	// Development development env value
	Development = Env("development")

	// Test test env value
	Test = Env("test")
)

// IsProduction is production
func (e Env) IsProduction() bool {
	if len(e) == 0 {
		return false
	}
	return e == Production
}

// IsDevelopment is development, nil is true
func (e Env) IsDevelopment() bool {
	if len(e) == 0 {
		return true
	}
	return e == Development
}

// IsTest is test
func (e Env) IsTest() bool {
	if len(e) == 0 {
		return false
	}
	return e == Test
}

// String the string form
func (e Env) String() string {
	return string(e)
}
