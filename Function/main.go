
// aturan fuction 1. Tidak bisa diawali dengan angka
// 2. Tidak boleh mengandung spasi
// 3. Tidak boleh mengandung simbol kecilai dengan(_)
// 4. Apibila func diawali huruf besar dapat dipanggil diluar package main, kalu kecil hanya dalam package main
package main
import "fmt"

func greet(name string)(text string){
	//text "Hello" + name
	//return
	return "halo: " + name

}
func add(x,y int)int{
	return x + y
}
func add1 (x,y int)(addres int, as int){
	//return x+y, x - y
	addres = x + y
	as = x + y
	return

}

func main (){
	fmt.Println(add1(9,8))
}