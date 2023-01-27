package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	meuPreço = 10.42
	meuNome  = "Product 1"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct(meuNome, meuPreço)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.Name)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.Price)
	assert.Equal(t, meuNome, p.Name)
	assert.Equal(t, meuPreço, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", meuPreço)

	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct(meuNome, 0)

	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct(meuNome, -1)

	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct(meuNome, meuPreço)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
