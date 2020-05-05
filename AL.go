// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
)

	var lineas [15]string
	var j int=0
	var opc int=0


func main() {

	identificadores := make([]string, 0)
	reservadas := make([]string, 0)
	numeros := make([]string, 0)
	operadores := make([]string, 0)
	simbolos := make([]string, 0)
	nodefinidos := make([]string, 0)

	fmt.Println("Introduce 1 o 2 para leer un codigo en basic")
	fmt.Scanf("%d\n", &opc)
	if(opc==1){
		archivo, error:= os.Open("./basic.txt")
		if error!= nil{
			fmt.Println("error, archivo no existe")
		}
		scanner:= bufio.NewScanner(archivo)
		for scanner.Scan(){
			linea:=scanner.Text()
			lineas[j]=linea
			j++
		}
	}else if(opc==2){
		archivo, error:= os.Open("./basic1.txt")
		if error!= nil{
			fmt.Println("error, archivo no existe")
		}
		scanner:= bufio.NewScanner(archivo)
		for scanner.Scan(){
			linea:=scanner.Text()
			lineas[j]=linea
			j++
		}
	}
	j=0
	for j := 0; j < 15; j++ {
		fmt.Println(lineas[j])
	}
	j=0

	for index, line := range lineas {
		line = lineas[index]
	    var words = strings.Split(line, " ")
	    for _, word := range words {
				if isReservedWord(word) {
					reservadas = append(reservadas, word)
				} else if isNumber(word) {
					numeros = append(numeros, word)
				} else if isOperator(word) {
					operadores = append(operadores, word)
				} else if isIdentifier(word) {
	      			identificadores = append(identificadores, word)
				} else if isSymbol(word) {
					simbolos = append(simbolos, word)
				} else{
					nodefinidos = append(nodefinidos, word)
				}
	    }
  	}

	fmt.Println("*****************************************")
  	fmt.Println("*\t\tTOKENS\t\t\t*")
  	fmt.Println("*****************************************")

	fmt.Println("\nIdentificadores: ", identificadores)
	fmt.Println("\nPalabras Reservadas: ", reservadas)
	fmt.Println("\nNumeros: ", numeros)
	fmt.Println("\nDelimitadores: ", simbolos)
	fmt.Println("\nOperadores:  ", operadores)
	fmt.Println("\nIndefinidos: ", nodefinidos)
	if(checarBegin(reservadas)==true){
		fmt.Println("\nErrores: El programa debe iniciar con BEGIN")
	}
	if(checarEnd(reservadas)==true){
		fmt.Println("\nErrores: El programa debe terminar con END")
	}
	fmt.Println("\n****************************************")
}

func isNumber (token string) bool {
	var re = regexp.MustCompile("^(\\d+|\\d*\\.\\d+)$")
	if(re.MatchString(token)){
		return true
	}
	return false
}

func isOperator(token string) bool {
	var re = regexp.MustCompile("[*|/|+|-|=]$")
	if(re.MatchString(token)){
		return true
	}
	return false
}

func isIdentifier(token string) bool {
	var re = regexp.MustCompile("^[A-Za-z]*[%$&#]$")
	if(re.MatchString(token)){
		return true
	}
	return false
}

func isSymbol(token string) bool {
	var re = regexp.MustCompile("[/|\\|(|)|?|;|{|}]$")
	if(re.MatchString(token)){
		return true
	}
	return false
}

func isReservedWord(token string) bool {

	var palabrasReservadas = [] string {"PRINT","LET","LIST","RUN","INPUT","GOTO","IF",
	"THEN","ELSE","STOP","END", "BEGIN","EDIT","AUTO","WHILE","FOR","NEXT","TO","LEN", "OR"}

	for _, reservada := range palabrasReservadas {
		if(strings.ToLower(reservada) == strings.ToLower(token)){
			return true
		}
	}
	return false;
}

func checarBegin(reservadas []string) bool {
	if(strings.ToLower(reservadas[0])!="begin"){
		return true
	}else{
		return false
	}
}

func checarEnd(reservadas []string) bool {
	if(strings.ToLower(reservadas[len(reservadas)-1])!="end"){
		return true
	}
	return false
}