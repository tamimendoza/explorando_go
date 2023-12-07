package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	p := Persona{"Juan", 1980}
	fmt.Println(p.Edad())
	p.SetAnioNacimiento(1985)
	fmt.Println(p.Edad())
	fmt.Println(p)
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type MiError struct {
	Cuando time.Time
	Que    string
}

func (e *MiError) Error() string {
	return fmt.Sprintf("Fecha %v, %s", e.Cuando, e.Que)
}

func run() error {
	return &MiError{
		time.Now(),
		"no funcionó",
	}
}

type Persona struct {
	Nombre         string
	AnioNacimiento int
}

func (p Persona) Edad() int {
	return time.Now().Year() - p.AnioNacimiento
}

func (p Persona) String() string {
	return fmt.Sprintf("Nombre: %s, Año de nacimiento: %d", p.Nombre, p.AnioNacimiento)
}

func (p *Persona) SetAnioNacimiento(anio int) {
	p.AnioNacimiento = anio
}

func main_funciones() {
	fmt.Println(suma(2, 3))
	a, b := intercambio("hola", "mundo")
	fmt.Println(a, b)
	fmt.Println(retorno(10))

	fnSumador := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(fnSumador(1, 2))
	fmt.Println(ejecutar(fnSumador))
	fmt.Println(ejecutar(math.Pow))

	pos, neg := sumador(), sumador()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func sumador() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func ejecutar(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func retorno(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func intercambio(a, b string) (string, string) {
	return b, a
}

func suma(a int, b int) int {
	return a + b
}

func main_struc() {
	v := Vector{1, 2}
	v.X = 4
	fmt.Println(v)

	vector := Vector{1, 2}
	p := &vector
	p.X = 9
	fmt.Println(vector)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(s)

	m = make(map[string]Punto)
	m["Bell Labs"] = Punto{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

type Punto struct {
	Lat, Long float64
}

var m map[string]Punto

type Vector struct {
	X, Y int
}

func main_punteros() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p - 1
	fmt.Println(j)
}

func main_defer() {
	defer fmt.Println("mundo")
	fmt.Println("hola")

	fmt.Println("contando")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("hecho")
}

func main_for() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	nombres := []string{"Juan", "Pedro", "Luis"}
	for indice, nombre := range nombres {
		fmt.Printf("El nombre %q está en el índice %d del slice\n", nombre, indice)
	}

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	contador := 0
	for contador < 10 {
		fmt.Println(contador)
		contador++
	}

	/* bucle infinito
	for {
	}
	*/

	fmt.Println("**************************")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}

func main_switch() {
	dia := "martes"

	switch dia {
	case "lunes":
		fmt.Println("Hoy es lunes")
	case "martes":
		fmt.Println("Hoy es martes")
	default:
		fmt.Println("No es ni lunes ni martes")
	}

	hora := time.Now()
	switch {
	case hora.Hour() < 12:
		fmt.Println("Buenos días!")
	case hora.Hour() < 17:
		fmt.Println("Buenas tardes.")
	default:
		fmt.Println("Buenas noches.")
	}

	evaluador(21)
	evaluador("Hola")
	evaluador(true)
}

func evaluador(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("El doble de %v es %v\n", v, v*2)
	case string:
		fmt.Printf("%q tiene %v bytes\n", v, len(v))
	default:
		fmt.Printf("No sé nada sobre el tipo %T!\n", v)
	}
}

func main_if() {
	edad := 30

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
}

func conversion_tipo() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(f)
	var z uint = uint(f)
	fmt.Println(x, y, z)

	variable := 42
	fmt.Printf("variable es de tipo %T\n", variable)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	nombres := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(nombres)

	a := nombres[0:2]
	b := nombres[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(nombres)

	// rangos
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	s1 := pow[2:4]
	fmt.Println(s1)

	s2 := pow[:3] // 0..2
	fmt.Println(s2)

	s3 := pow[:]
	fmt.Println(s3)

	s4 := pow[4:]
	fmt.Println(s4)

	// slices son nulos
	var s5 []int
	fmt.Println(s5, len(s5), cap(s5))
	if s5 == nil {
		fmt.Println("nulo!")
	}
	s5 = append(s5, 2)
	fmt.Println(s5, len(s5), cap(s5))

	// creando slices con make
	s6 := make([]int, 5)
	fmt.Println(s6, len(s6), cap(s6))

	s7 := make([]int, 0, 5)
	fmt.Println(s7, len(s7), cap(s7))

	s8 := s7[:2]
	fmt.Println(s8, len(s8), cap(s8))

	// recorrer un slice
	var aleatorio = []int{2, 3, 5, 7, 11, 13}
	for i, v := range aleatorio {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// slices de slices
	tablero := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	tablero[0][0] = "X"
	tablero[2][2] = "O"
	tablero[1][2] = "X"
	tablero[1][0] = "O"
	tablero[0][2] = "X"
	for i := 0; i < len(tablero); i++ {
		fmt.Printf("%s\n", strings.Join(tablero[i], " "))
	}
}

func variables() {
	// declarar variables
	var edad1 int
	edad1 = 30
	fmt.Println("edad1:", edad1)

	var edad2 int = 40
	fmt.Println("edad2:", edad2)

	var i, j int = 50, 60
	fmt.Println("i:", i, "j:", j)

	edad3 := 70
	fmt.Println("edad3:", edad3)

	var (
		nombre  string = "Juan"
		edad4   int    = 30
		soltero bool   = false
	)
	fmt.Println("nombre:", nombre, "edad4:", edad4, "soltero:", soltero)

	var saldo float32
	fmt.Println("saldo:", saldo)

	const pi = 3.1416
	fmt.Println("pi:", pi)
}

func arreglo() {
	var arr [2]string
	arr[0] = "Hola"
	arr[1] = "Mundo"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
