package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/sallescosta/goexpert/api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	meuNome  = "Product 1"
	meuPreço = 10.34
	perPage  = 10
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct(meuNome, meuPreço)
	assert.Nil(t, err)

	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	var precoSorteado = rand.Float64() * 100

	for i := 1; i < 33; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Produto %d", i), precoSorteado)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, perPage, "asc")
	assert.NoError(t, err)

	assert.Len(t, products, perPage)
	assert.Equal(t, "Produto 1", products[0].Name)
	assert.Equal(t, "Produto 10", products[9].Name)

	products, err = productDB.FindAll(2, perPage, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, perPage)
	assert.Equal(t, "Produto 11", products[0].Name)
	assert.Equal(t, "Produto 20", products[9].Name)

	products, err = productDB.FindAll(3, perPage, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, perPage)
	assert.Equal(t, "Produto 21", products[0].Name)
	assert.Equal(t, "Produto 30", products[9].Name)

	products, err = productDB.FindAll(4, perPage, "asc")
	println(products)
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Produto 31", products[0].Name)
	assert.Equal(t, "Produto 32", products[1].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct(meuNome, meuPreço)
	assert.NoError(t, err)

	db.Create(product)
	productDB := NewProduct(db)

	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, meuNome, product.Name)
	assert.Equal(t, meuPreço, product.Price)

}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct(meuNome, meuPreço)
	assert.NoError(t, err)

	db.Create(product)
	productDB := NewProduct(db)

	product.Name = "Produto 2"
	err = productDB.Update(product)
	assert.NoError(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Produto 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct(meuNome, meuPreço)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
