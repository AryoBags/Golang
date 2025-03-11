
// Interface bisa digunakan sebagai penent fungsi fungsi data kosong atau yag harus di implemntasi
// Implementasi menggunakan method
// method bisa digunakan oleh dua hal yaitu type alias dan struct
// type alias adalah tipe data baru yang kita ambil dari tipe lain
// type struct adalah kumpulan beberapa variabel menjadi sebuah objek yang bisa kita gunakan berkali-kali
package main
import "fmt"

// Alias
type CentiMeters float64

func  (c CentiMeters) IsToooLong(){
	if c > 10{
		fmt.Printf("panjang terlalu panjang\n")
	}else{
		fmt.Printf("just Fine\n")
	}
	
}
// struct
type rectangle struct{

	width float64
	height float64
}
type shape interface{
	getArea()	float64
	getPerime() float64
}
func (r rectangle) getArea() float64 {
	return r.width * r.height


}
func (r rectangle) getPerime() float64 {
	return 2*r.width + 2*r.height
}
func measureShape(s shape){
	fmt.Printf("area: %.2f\n", s.getArea())
	fmt.Printf("perimeter: %.2f\n", s.getPerime())
}

type circle struct{
	radius float64

}
func (c circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
	}
func (c circle) getPerime() float64 {
	return 2 * 3.14 * c.radius
}

type square struct{
	side float64
}
func (s square) getArea() float64 {
	return s.side * s.side
}
func (s square) getPerime() float64 {

		return 4 * s.side
}


func main(){
	myvar := CentiMeters(13)
	fmt.Printf("Type: %T, value %v\n", myvar, myvar)
	myvar.IsToooLong()

	myRetacle := rectangle{
		width: 10.0,
		height: 20.0,
	}
	myCir := circle{

		radius: 4,
	}
	mySque := square{
		side: 10.0,
	}
	fmt.Printf("Type: %T, value %v\n", myRetacle, myRetacle)
	fmt.Printf("Type: %T, value %v\n", myCir, myCir)
	fmt.Printf("Type: %T, value %v\n", mySque, mySque)
	measureShape(myRetacle)
	measureShape(myCir)
	measureShape(mySque)
	
}