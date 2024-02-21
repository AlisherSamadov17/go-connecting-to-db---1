package city

import (
	"fmt"

	"github.com/google/uuid"
)

type City struct {
 Id uuid.UUID `json:"id"`
 Name string `json:"name"`
 Population int `json:"population"`
 Country_id uuid.UUID `json:"country_id"`
 Created_at string `json:"created_at"`
 Updated_at string `json:"updated_at"`
}

func (c City) NewCity() City {
	var (
	  id string
	  name string
	  population int
	  country_id string
	)
  
	fmt.Print("Id: ")
	fmt.Scan(&id)
	fmt.Print("name: ")
	fmt.Scan(&name)
	fmt.Print("population: ")
	fmt.Scan(&population)
	fmt.Print("country_id: ")
	fmt.Scan(&country_id)
  
	return City{
	  Id: uuid.MustParse(id),
	  Name: name,
	  Population: population,
	  Country_id: uuid.MustParse(country_id),  
	}
  }
  