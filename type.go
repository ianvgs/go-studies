package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

// Tipo uma classe
type Car struct {
	Name  string
	Model string
	Price float64
}

var cars []Car

func createCars() {
	cars = append(cars, Car{Name: "Ferrari", Price: 100, Model: "Sei nao"})
	cars = append(cars, Car{Name: "Mercedes", Price: 200, Model: "Sei nao"})
	cars = append(cars, Car{Name: "Porsche", Price: 300, Model: "Sei nao"})
}

//Tipo os métodos de uma classe
/* func (c Car) Andar() {
	fmt.Println("O carro", c.Name, "está andando")
}
*/
func main() {
	/* 	carro := Car{"fusca", "wv"}
	   	carro.Model = "Audi"
	   	fmt.Println(carro.Name)
	   	carro.Andar() */

	createCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createC)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createC(c echo.Context) error {
	//isso é só o formato
	car := new(Car)

	// compara o structc com o objeto recebido
	if err := c.Bind(car); err != nil {
		return err
	}
	////////////////////* é o ponteiro, local da memoria pra car append

	cars = append(cars, *car)

	saveCar(*car)

	return c.JSON(200, cars)
}

func saveCar(car Car) error {

	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}

	//fecha conexão depois de fazer oq precisa
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, price) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	return nil

	//substitui res por _ pra compilar, nao da em todas
	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil

}

//sqlite3 cars.db
//create table cars (name text, price number, model text);
//.table

//sqlite3 cars.db
//select * from cars
