package pr1libraries

import (
	"fmt"
	"reflect"
	s "strings"
)

func ByteRune() {
	var myByte byte = 'a'
	var myRune rune = '♥'

	fmt.Println("\n============ ByteRune =============\n")
	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
	fmt.Println("\n==============================\n")
}

func UintFloat() {
	var (
		angka1     uint8   = 21
		angka2     uint8   = 17
		angkaFloat float64 = 7.1
	)

	fmt.Println("\n============ UintFloat =============\n")
	// ./typecast.go:14:11: cannot use "abc" (type string) as type uint8 in assignment
	//angka1 = "abc"
	fmt.Println(angka1 + angka2)
	// ./typecast.go:15:21: invalid operation: angka1 + angkaFloat (mismatched types uint8 and float64)
	//fmt.Println(angka1 + angkaFloat)
	fmt.Println(float64(angka1) + angkaFloat)
	fmt.Println("\n==============================\n")
}

func String() {
	// Definisi string
	var str1 string = "UGM"
	var str2 = "Yogyakarta"
	var str3 = "universitas\ngadjah mada"

	var str3backtick = `universitas\ngadjah mada`

	// error: illegal rune literal
	//var str3singlequoted = 'universitas gadjah mada'

	fmt.Println("\n============ String =============\n")
	fmt.Println(str1)
	fmt.Println(len(str1))
	fmt.Println(s.Contains(str1, "GM"))
	fmt.Println(s.Title(str3))
	fmt.Println(str1[0])
	fmt.Println(s.Join([]string{str1, str2}, " "))
	fmt.Println(s.Join([]string{str3, str2}, "\n"))
	fmt.Println(s.Join([]string{str3backtick, str2}, "\n"))
	fmt.Println(reflect.TypeOf(str1))
	fmt.Println(reflect.TypeOf(str2))
	fmt.Println()
	fmt.Println("\n==============================\n")
}

func Boolean() {
	var (
		hasilPerbandingan bool
		angka1            uint8 = 21
		angka2            uint8 = 17
	)

	fmt.Println("\n============ Boolean =============\n")
	hasilPerbandingan = angka1 < angka2
	fmt.Printf("angka1 = %d\n", angka1)
	fmt.Printf("angka2 = %d\n", angka2)
	fmt.Println(reflect.TypeOf(hasilPerbandingan))
	fmt.Println(hasilPerbandingan)
	fmt.Println("\n==============================\n")
}

func DefaultVariable() {
	// unsigned-integer
	var defUint8 uint8
	var defUint16 uint16
	var defUint32 uint32
	var defUint64 uint64
	var defUint uint

	// signed-integer
	var defInt8 int8
	var defInt16 int16
	var defInt32 int32
	var defInt64 int64
	var defInt int

	// string
	var defString string

	// floating-point
	var defFloat32 float32
	var defFloat64 float64

	// complex
	var defComplex64 complex64
	var defComplex128 complex128

	// alias
	var defByte byte
	var defRune rune

	fmt.Println("\n============ Default Variable =============\n")
	fmt.Println("\nNilai default untuk uint8 = ", defUint8)
	fmt.Println("Nilai default untuk uint16 = ", defUint16)
	fmt.Println("Nilai default untuk uint32 = ", defUint32)
	fmt.Println("Nilai default untuk uint64 = ", defUint64)
	fmt.Println("Nilai default untuk uint = ", defUint)

	fmt.Println("\nNilai default untuk int8 = ", defInt8)
	fmt.Println("Nilai default untuk int16 = ", defInt16)
	fmt.Println("Nilai default untuk int32 = ", defInt32)
	fmt.Println("Nilai default untuk int63 = ", defInt64)
	fmt.Println("Nilai default untuk int = ", defInt)

	fmt.Println("\nNilai default untuk string = ", defString)

	fmt.Println("\nNilai default untuk float32 = ", defFloat32)
	fmt.Println("Nilai default untuk float64 = ", defFloat64)

	fmt.Println("\nNilai default untuk complex64 = ", defComplex64)
	fmt.Println("Nilai default untuk complex128 = ", defComplex128)

	fmt.Println("\nNilai default untuk byte = ", defByte)
	fmt.Println("Nilai default untuk rune = ", defRune)
	fmt.Println("\n==============================\n")
}

func Constanta() {
	const mainCodingLang = "Go"
	const kiamatMakinDekat = true

	const angka1, angka2 = 25, 8

	const (
		nomorPegawai = "P001"
		gaji         = 50000000
	)

	const negaraKu string = "Indonesia"

	const gajiBersihSetelahSetorIstri = gaji - 49000000

	fmt.Println("\n============ Konstanta =============\n")
	fmt.Println(mainCodingLang)
	fmt.Println(kiamatMakinDekat)
	fmt.Println(angka1)
	fmt.Println(angka2)
	fmt.Println(nomorPegawai)
	fmt.Println(gaji)
	fmt.Println(negaraKu)
	fmt.Println(gajiBersihSetelahSetorIstri)
	fmt.Println("\n==============================\n")
}

func Pointer() {
	i, j := 42, 2701

	// p berisi memory address dari i
	p := &i
	// tampilkan isi dari memory address p
	fmt.Println(*p)
	// isi dari memory address yang ditunjuk p diubah
	*p = 21
	// implikasinya pada variabel i:
	fmt.Println(i)

	// p berisi memory address dari j
	p = &j
	// isi dari memory address yang ditunjuk p, diubah
	// menjadi isi memoery address yang lama, dibagi 37
	*p = *p / 37
	// implikasinya pada variabel j
	fmt.Println(j) // see the new value of j

	var pa *int

	fmt.Println("\n============ Konstanta =============\n")
	fmt.Printf("pointer pa dengan tipe %T dan nilai %v\n", pa, pa)
	fmt.Println("\n==============================\n")
}

func ElseIf() {
	var a int = 200

	fmt.Println("\n============ If Else =============\n")
	if a == 10 {
		fmt.Printf("Nilai a = 10\n")
	} else if a == 20 {
		fmt.Printf("Nilai a = 20\n")
	} else if a == 30 {
		fmt.Printf("Nilai a = 30\n")
	} else {
		fmt.Printf("Semua nilai salah\n")
	}
	fmt.Printf("Nilai dari a adalah: %d\n", a)
	fmt.Println("\n==============================\n")
}

func Switch() {
	var nilaiAngka int = 20
	var nilaiHuruf string

	switch nilaiAngka {
	case 90:
		nilaiHuruf = "A"
	case 80:
		nilaiHuruf = "B"
	case 50, 60, 70:
		nilaiHuruf = "C"
	default:
		nilaiHuruf = "D"
	}

	fmt.Println("\n============ Switch =============\n")
	switch {
	case nilaiHuruf == "A":
		fmt.Println("Apik tenan!")
	case nilaiHuruf == "B":
		fmt.Println("Lumayan lah")
	case nilaiHuruf == "C", nilaiHuruf == "D":
		fmt.Println("Lulus sih ... tapi ..")
	case nilaiHuruf == "E":
		fmt.Println("Nangis bombay")
	default:
		fmt.Println("Nilai gak jelas, seperti wajah dosennya!")
	}
	fmt.Printf("Nilai anda =  %s\n", nilaiHuruf)
	fmt.Println("\n==============================\n")
}

func Looping() {
	var a int = 10

	fmt.Println("\n============ Looping =============\n")
	for a < 20 {
		if a == 12 {
			a += 1
			continue
		}
		a++
		if a > 15 {
			break
		}
		fmt.Printf("Nilai a: %d\n", a)
	}
	fmt.Println("\n==============================\n")
}

func Defer() {
	fmt.Println("\n============ Defer =============\n")
	defer fmt.Println("world")

	fmt.Println("hello")
	fmt.Println("\n==============================\n")
}
