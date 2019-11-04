package main
import "fmt"

type Stringer interface{
	//This is the method in this interface
	String() string
	Area() int
}

type Point struct{x, y int}

func PointToString(p Point) string{
	return fmt.Sprintf("Point{%d, %d}", p.x,p.y)
}

func(p Point) String() string{
	return fmt.Sprintf("Point{%d, %d}", p.x,p.y)
}

func(p Point) Area() int {
	return p.x *  p.y
}

func main() {
	var v Stringer
	var corner = Point{1,2}
	var boiling = Celsius(100)

	v = corner
	fmt.Println(v.String())
	fmt.Println(v)
	fmt.Println(v.String())
	fmt.Println(v.Area())

	v = boiling.ToFahrenheit()
	fmt.Println(v)
	fmt.Println(v.String())
	fmt.Println(v.Area())
}


type Celsius float32
type Fahrenheit float32

func (t Celsius) String() string {return fmt.Sprintf("%g°C", t)}
func (t Fahrenheit) String() string {return fmt.Sprintf("%g°F", t)}
func (t Celsius) ToFahrenheit() Fahrenheit {return Fahrenheit((t*9/5+32))}
func (t Fahrenheit) Area() int {return -1}
