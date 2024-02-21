package country

import (
	"fmt"

	"github.com/google/uuid"
)

type Country struct {
	Id uuid.UUID    `json:"id"`
	Name string     `json:"name"`
	Currency string `json:"currency"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func (c Country) NewCountry() Country {
  var (
	id string
	name string
	currency string
  )

  fmt.Print("Id: ")
  fmt.Scan(&id)
  fmt.Print("name: ")
  fmt.Scan(&name)
  fmt.Print("currency: ")
  fmt.Scan(&currency)


  return Country{
	Id: uuid.MustParse(id),
	Name: name,
	Currency: currency,
  }
}

  











































































































