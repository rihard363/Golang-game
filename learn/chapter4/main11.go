package main
import ("fmt"; "math")
type Circle struct {
    x float32
    y float32
    r float32
}
type Rectangle struct {
    x1, y1, x2, y2 float32
}
type geometry interface {
  perimetr() float32
  area() float32 
}
func (l Rectangle) area() float32 {
			return (l.x2-l.x1)*(l.y2-l.y1)
}
func (l Rectangle) perimetr() float32 {
			return 2*(l.x2-l.x1)+2*(l.y2-l.y1)
}
func (c Circle) area() float32 {
			return math.Pi*c.r*c.r
}
func (c Circle) perimetr() float32 {
			return math.Pi*2*c.r
}
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perimetr())
}
func main() {
	c := Circle{0, 0, 5}
	l := Rectangle{0, 0, 10, 10}
	measure(c)
	measure(l)
}