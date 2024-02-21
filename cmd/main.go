package main

import (
	"database/sql"
	"fmt"
	"main/city"
	"main/country"
	"main/inventory"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var (
	Country int = 1
	City int = 2
)

func main()  {
// ---------------------------------------------
	db,err :=connectDb()
	if err != nil {
		fmt.Println("error while connecting to main and opening db file",err)
	   return
	}
	defer db.Close()
// ---------------------------------------------
 inv := inventory.Inventory{Db: db}
 fmt.Println("------------------------------------")

 fmt.Println("1)Country")
 fmt.Println("2)City")
fmt.Println("------------------------------------")
 var numberFromInput int
 fmt.Println("Enter your choice:")
 fmt.Scan(&numberFromInput)

switch numberFromInput {
case Country:
	fmt.Println("1-adding a COUNTRY")
	fmt.Println("2-getting a COUNTRY by id")
	fmt.Println("3-Country GetAll")
	fmt.Println("4-updating Country")
	fmt.Println("5-deleting Country")
	fmt.Println("---------------------------------")
	fmt.Println()
	fmt.Print("Enter your choice: ")
	var inputForUser int;
	fmt.Scan(&inputForUser)
	switch inputForUser {
	case 1:
		c := country.Country{}
		if err := inv.CreateCountry(c.NewCountry());err != nil{
			fmt.Println("error while adding Country",err)
		return
		}
	fmt.Println("Country successfully added to db !")
case 2:
	var id string
	fmt.Print("enter the id :")
	fmt.Scan(&id)

	s,err :=inv.GetByIDCountry(id)
	if err != nil {
		fmt.Println("error while country by id:",err)
		return
	}
	fmt.Println(s)
case 3:
	c,err :=inv.GetAllCountry()
	if err != nil{
		fmt.Println("error while gettin all countries",err)
		panic(err)
	}
	fmt.Println(c)

case 4:
	var id string
	var name string
	var	currency string

	fmt.Print("id: ")
	fmt.Scan(&id)

	fmt.Print("name: ")
	fmt.Scan(&name)

	fmt.Print("currency: ")
	fmt.Scan(&currency)

	coun := country.Country{
		Id: uuid.MustParse(id),
		Name: name,
		Currency: currency,
	}
	if err := inv.UpdateCountry(coun);err != nil {
       panic(err)
	}
	fmt.Println("country updated !")
case 5:
	var id string
	fmt.Print("id: ")
	fmt.Scan(&id)

	if err := inv.DeleteCountry(uuid.MustParse(id));err != nil{
		fmt.Println("error while deleting",err)
		panic(err)
	}
	fmt.Println("country deleted !")

default:
	fmt.Println("error while entering the commands")
	} 
//----------------------------------------------------------- 
case City:
	fmt.Println("1-adding a CITY")
	fmt.Println("2-getting a CITY by id")
	fmt.Println("3-City GetAll")
	fmt.Println("4-updating City")
	fmt.Println("5-deleting City")
	fmt.Println("---------------------------------")
	fmt.Println()
	fmt.Print("Enter your choice: ")
	var inputForUser int;
	fmt.Scan(&inputForUser)
	switch inputForUser{
	case 1:
		c := city.City{}
		if err := inv.CreateCity(c.NewCity());err != nil{
			fmt.Println("error while adding city",err)
			return
		}
	fmt.Println("City successfully added to db !")
	case 2:
		var id string
		fmt.Print("enter the id :")
		fmt.Scan(&id)

		s,err := inv.GetByIDCity(id)
		if err != nil{
			fmt.Println("error while city by id",err)
			return
		}
		fmt.Println(s)
	case 3:
		c,err :=inv.GetAllCity()
		if err != nil{
			fmt.Println("error while getting all cities",err)
		  panic(err)
		}	
		fmt.Println(c)
	case 4:
     var id string
	 var name string
	 var population int

	 fmt.Print("id:")
	 fmt.Scan(&id)

	 fmt.Print("name:")
	 fmt.Scan(&name)

	 fmt.Print("population:")
	 fmt.Scan(&population)

	 cit := city.City{
		Id: uuid.MustParse(id),
		Name: name,
		Population:population,
	 }
	 if err := inv.UpdateCity(cit);err != nil{
		panic(err)
	 }
	 fmt.Println("country updated !")
	case 5:
		var id string
		fmt.Print("id:")
		fmt.Scan(&id)

      if err := inv.DeleteCity(uuid.MustParse(id));err != nil{
		fmt.Println("error while deleting",err)
		panic(err)
	  }
	  fmt.Println("country deleted !")

	default:
		fmt.Println("error while entering the commands")
	}
}
}

func connectDb() (*sql.DB,error) {
	connecting :="host=localhost port=5432 user=person password=1234 database=useanduse sslmode=disable"
    db,err :=sql.Open("postgres",connecting)
	if err != nil {
		fmt.Println("error while opening db file",err)
	}
	return db,nil
}