package main
var a int
/*
	123.123123:float64
*/
/*
	false:bool
*/
var ToBe bool = false
var c float64 = 123.123123
var d []int
type intAlias int
func main() {
	/*
		"Hello World":string
	*/
	print("Hello World")
	/*
		"dadas	":string
	*/
	/*
		`
        		`:string	*/
	/*
		'	':rune	*/
	a, v, h := "dadas	", `
        		`, '	';
	var t struct{
			b int
			nb intAlias
		}

	/*
		0:int
	*/
	/*
  		i<10: bool
		i:int
		10:int
	*/
  	/*
		i:int
	*/
	for i := 0;i<10;i-- {
		/*
			a:string
		*/
		print(a)
	}
	/*
		3:int
	*/
	/*
		tag:int
	*/
	switch tag := 3;tag{
		/*
			0:int
		*/
		/*
			1:int
		*/
		/*
			2:int
		*/
		/*
			3:int
		*/
		case 0, 1, 2, 3:
		/*
			4:int
		*/
		/*
			5:int
		*/
		/*
			6:int
		*/
		/*
			7:int
		*/
		case 4, 5, 6, 7:
		default:
	}
}
