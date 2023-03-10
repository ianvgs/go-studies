/* package main

func main(){
	println("hello wolrd")
} */

package main

import "fmt"

/* import "net/http" */

/* func main() {
		http.HandleFunction("/", func(w http.ResponseWriter, r *http.Request))
	 http.ListenAndServe(":8080", nil)

	// := cria e ja atribui valor
	nome := "Ian"
	nome = "novo nome"
	fmt.Println(nome)

	var a int = 10
	a = 11
	fmt.Println(a)

}
*/

func main() {
	/* result, err := soma(1, 2)
	fmt.Println(result, err) */
	//printa 3 <nil>

	result, err := soma(1, 20)

	//printa 0 Soma maior que 10

	//verifica se tem erro e trata
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result, err)
}

//tem que colocar o retorno no ()type senão não pode ter return na função
func some(a int, b int) int {
	if a+b > 10 {
		return 0
	}
	return a + b
}

//tem que colocar o retorno no ()type senão não pode ter return na função e pode ter dois retornos
func soma(a int, b int) (int, error) {

	//maior que 10 retorna 0 e lança erro
	if a+b > 10 {
		return 0, fmt.Errorf("Soma maior que 10")
	}
	//retorna a soma e o erro em branco pq declarou que ia retornar erro
	return a + b, nil
}
