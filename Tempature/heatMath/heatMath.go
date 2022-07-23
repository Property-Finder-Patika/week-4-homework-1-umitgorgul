package heatMath

//interface zone
type HeatFunction interface {
	Calculate(arg float64, toChange string) float64 
	GetName() string
}

// Celsius zone
type Celsius struct {
	Name string
}

func (c Celsius) GetName() string {
	return c.Name
}

func (c Celsius) Calculate(arg float64, toChange string) float64 {
	if toChange == "Fahrenheit" { // change to fahrenheit
		return arg*9/5 + 32

	} else { // change to kelvin
		return arg - 273.15
	}
}

// Fahrenheit zone
type Fahrenheit struct {
	Name string
}

func (f Fahrenheit) GetName() string {
	return f.Name
}

func (f Fahrenheit) Calculate(arg float64, toChange string) float64 {
	if toChange == "Kelvin" { // change to fahrenheit
		return (arg + 459.67) * (5 / 9)

	} else { // change to celcius
		return (arg - 32) * (5 / 9)
	}
}

// Kelvin zone
type Kelvin struct {
	Name string
}

func (k Kelvin) GetName() string {
	return k.Name
}

func (k Kelvin) Calculate(arg float64, toChange string) float64 {
	if toChange == "Fahrenheit" { // change to fahrenheit
		return arg*9/5 + 32

	} else { // change to celcius
		return arg + 273.15
	}

}
