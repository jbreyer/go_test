package tempconv

type Celsius float64
type Fahrenheit float64

const(
	AbsoluteZeroC = -273.15
	FreezingC = 0
	BoilingC = 100
)

func CToF(celsius Celsius) Fahrenheit {return Fahrenheit(9/5*celsius+32)}
func FToC(fahrenheit Fahrenheit) Celsius {return Celsius((fahrenheit-32)*5/9)}
