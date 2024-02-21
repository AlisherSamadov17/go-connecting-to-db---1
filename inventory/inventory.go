package inventory

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"main/city"
	"main/country"

	"github.com/google/uuid"
)

type Inventory struct{
	Db *sql.DB
}

func (i Inventory) New(db *sql.DB) Inventory {
	return Inventory{
		Db: db,
	}
}

// ---------_countries_---------
func (i Inventory) CreateCountry(c country.Country) error {
	id := uuid.NewString()
	_,err := i.Db.Exec(`insert into countries(id,name,currency,updated_at) values ($1,$2,$3,NOW())`,id,c.Name,c.Currency)
    if err != nil {
		fmt.Println("error while getting from db",err)
		return err
	} 
	return nil
}

func (i Inventory) GetByIDCountry(id string) (country.Country,error)  {
	cs := country.Country{}
	row :=i.Db.QueryRow(`select * from countries where id = $1`,id)
	if err := row.Scan(&cs.Id,&cs.Name,&cs.Currency,&cs.Created_at,&cs.Updated_at);err != nil {
		if !errors.Is(err,sql.ErrNoRows){
			return country.Country{},err
		}
	}
	return cs,nil
}

func (i Inventory) GetAllCountry() ([]country.Country,error) {
	cs := []country.Country{}
	rows,err := i.Db.Query(`select * from countries`)
	if err != nil {
		return nil,err
	}
	for rows.Next(){
		c :=country.Country{}
		if err = rows.Scan(&c.Id,&c.Name,&c.Currency,&c.Created_at,&c.Updated_at);err != nil {
			return nil,err
		}
		cs = append(cs, c)
	}
	return cs,nil
}

func (i Inventory) UpdateCountry(c country.Country) error {
	if _,err := i.Db.Exec(`update countries set name=$1,currency=$2,updated_at=NOW() where id = $3`,c.Name,c.Currency,c.Id);err != nil {
       log.Println("error while getting from db",err)
	   return err
	}
	return nil
}

func (i Inventory) DeleteCountry(id uuid.UUID) error {
	if _,err := i.Db.Exec(`delete from countries where id = $1`,id);err != nil {
		return err
	}
	return nil
}


//---------_cities_---------
func (i Inventory) CreateCity(c city.City) error {
	id := uuid.NewString()
	_,err := i.Db.Exec(`insert into cities(id,name,population,country_id,updated_at) values ($1,$2,$3,$4,NOW())`,id,c.Name,c.Population,c.Country_id)
    if err != nil {
		fmt.Println("error while getting from db",err)
		return err
	} 
	return nil
}

func (i Inventory) GetByIDCity(id string) (city.City,error)  {
	cs := city.City{}
	row :=i.Db.QueryRow(`select * from cities where id = $1`,id)
	if err := row.Scan(&cs.Id,&cs.Name,&cs.Population,&cs.Country_id,&cs.Created_at,&cs.Updated_at);err != nil {
		if !errors.Is(err,sql.ErrNoRows){
			return city.City{},err
		}
	}
	return cs,nil
}

func (i Inventory) GetAllCity() ([]city.City,error) {
	cs := []city.City{}
	rows,err := i.Db.Query(`select * from cities`)
	if err != nil {
		return nil,err
	}
	for rows.Next(){
		c :=city.City{}
		if err = rows.Scan(&c.Id,&c.Name,&c.Population,&c.Country_id,&c.Created_at,&c.Updated_at);err != nil {
			return nil,err
		}
		cs = append(cs, c)
	}
	return cs,nil
}

func (i Inventory) UpdateCity(c city.City) error {
	if _,err := i.Db.Exec(`update cities set name=$1,population=$2,updated_at=NOW() where id = $3`,c.Name,c.Population,c.Id);err != nil {
       log.Println("error while getting from db",err)
	   return err
	}
	return nil
}

func (i Inventory) DeleteCity(id uuid.UUID) error {
	if _,err := i.Db.Exec(`delete from cities where id = $1`,id);err != nil {
		return err
	}
	return nil
}