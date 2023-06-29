package models

import (
	"fmt"

	"github.com/pkg/errors"
)

type Product struct {
	Id       uint64
	Name     string
	Price    float64
	Quantity int
	Amount   float64
	Category Category
}

func (s Store) NewProduct(product Product) (bool, error) {
	sql := "insert into products (name, price, quantity, amount, category) values ($1, $2, $3, $4, $5)"
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.Quantity, product.Amount, product.Category.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s Store) GetProducts() ([]Product, error) {
	sql := `select c.id, c.description, 
          p.id, p.name, p.price, p.quantity, p.amount  
          from products as p 
          inner join category as c on c.id = p.category order by p.id asc`
	rs, err := s.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var products []Product
	for rs.Next() {
		var product Product
		err := rs.Scan(&product.Category.Id,
			&product.Category.Description,
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Amount)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (s Store) SearchProducts(search string) ([]Product, error) {
	search = fmt.Sprintf("%%%s%%", search)
	sql := `select c.id, c.description, 
          p.id, p.name, p.price, p.quantity, p.amount  
          from products as p 
          inner join category as c on c.id = p.category
          where c.description like $1 or p.name like $2`
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rs, err := stmt.Query(search, search)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var products []Product
	for rs.Next() {
		var product Product
		err := rs.Scan(&product.Category.Id,
			&product.Category.Description,
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Amount)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (s Store) GetProductById(id uint64) (Product, error) {
	sql := "select * from products where id = $1"
	rs, err := s.DB.Query(sql, id)
	if err != nil {
		return Product{}, err
	}
	defer rs.Close()
	var product Product
	if rs.Next() {
		err := rs.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Amount, &product.Category.Id)
		if err != nil {
			return Product{}, err
		}
	}
	return product, nil
}

func (s Store) UpdateProduct(product Product) (int64, error) {
	sql := "update products set name = $1, price = $2, quantity = $3, amount = $4, category = $5 where id = $6"
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(product.Name, product.Price, product.Quantity, product.Amount, product.Category.Id, product.Id)
	if err != nil {
		return 0, err
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s Store) DeleteProduct(id uint64) (int64, error) {
	sql := "delete from products where id = $1"
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (p *Product) PriceToString() string {
	return fmt.Sprintf("%.2f", p.Price)
}

func (p *Product) AmountToString() string {
	return fmt.Sprintf("%.2f", p.Amount)
}
func (s Store) CreateCat() {

	const qry = `create table if not exists category(
    id serial primary key,
    description varchar(100) not null
  );`

	// Exec executes a query without returning any rows.
	if _, err := s.DB.Exec(qry); err != nil {
		err = errors.Wrapf(err,
			"Events table creation query failed (%s)",
			qry)
		return
	}
	return
}
func (s Store) CreatePro() {

	const qry = `create table if not exists products(
    id bigserial primary key,
    name varchar(255) not null,
    price real not null,
    quantity integer default 0,
    amount real default 0.0,
    category bigint not null,
    constraint products_category_fk foreign key(category)
    references category(id)
   );`

	// Exec executes a query without returning any rows.
	if _, err := s.DB.Exec(qry); err != nil {
		err = errors.Wrapf(err,
			"Events table creation query failed (%s)",
			qry)
		return
	}
}
func (s Store) CreateUser() {

	const qry = `create table if not exists users(
    id bigserial primary key,
    firstname varchar(15) not null,
    lastname varchar(20) not null,
    email varchar(40) not null unique,
    password varchar(100) not null,
    status char(1) default '0'
  );`

	// Exec executes a query without returning any rows.
	if _, err := s.DB.Exec(qry); err != nil {
		err = errors.Wrapf(err,
			"Events table creation query failed (%s)",
			qry)
		return
	}
}
