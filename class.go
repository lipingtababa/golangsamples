package main
import "fmt"

type Point struct{x, y int}

func PointToString(p Point) string{
	return fmt.Sprintf("Point{%d, %d}", p.x,p.y)
}

func(p Point) String() string{
	return fmt.Sprintf("Point{%d, %d}", p.x,p.y)
}


func main() {
	p := Point{3,5}
	fmt.Println(PointToString(p))
	fmt.Println(p.String())
	fmt.Println(p)

	var t Celsius = 21
	fmt.Println(t.String())
	fmt.Println(t)
	fmt.Println(t.toFahrenheit())
}


type Celsius float32
type Fahrenheit float32

func (t Celsius) String() string {return fmt.Sprintf("%g°C", t)}
func (t Fahrenheit) String() string {return fmt.Sprintf("%g°F", t)}
func (t Celsius) toFahrenheit() Fahrenheit {return Fahrenheit((t*9/5+32))}
