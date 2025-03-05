// pointer adalah sebuah varialbel yang menyimpan memori addres
package main
	
import "fmt"
//pointer
func incrementPoint(var1 *int){
	*var1  = *var1 + 1
}
// copy variabel
func increment(var1 int){
	var1 = var1 + 1
}
func main(){
	var1 := 1
	fmt.Printf("my variable is %d\n", var1)
	increment(var1)
	fmt.Printf("my variable is %d\n", var1)
	incrementPoint(&var1)
	fmt.Printf("my klmlsk %x kdvariable is %d\n", var1,var1)
	//copy addres
	fmt.Printf("my klmlsk %x kdvariable is %d\n", &var1,var1)

	var var2 *int
	fmt.Printf("my variable is %d\n", var2)

	var2  = new(int)
	*var2 = 34
	fmt.Printf("my variable is %v\n", *var2)

	var2  = &var1
	*var2 = 341
	fmt.Printf("my variable is %v\n", var1)
}

// pointer dapat merubah variabel dari jauh dengan tanda (*)
// Simbol untuk mengirim addres dari variabel adalah tanda (&)
// pointer dapat digunakan untuk membuat MT value yang bisa keluar 0 atau nill
// Pointer dapat ditambah dengan data baru  dengan kata "new"
// dan dapat memanggil addres dari valriabel lain 