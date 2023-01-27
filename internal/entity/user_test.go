package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var name = "John Doe"
var email = "abc@abc.com"
var password = "abcde"

// essas variáveis à cima, são utilizadas no método NewUser() para serem comparadas nas assertions

func TestNewUser(t *testing.T) {
	user, err := NewUser(name, email, password)

	assert.Nil(t, err)                 // expect que o nao haja erro...   ou err = nil
	assert.NotNil(t, user)             // expect que o user não seja nil
	assert.NotEmpty(t, user.ID)        // expect user.ID não esteja em branco
	assert.NotEmpty(t, user.Password)  // expect user.Password não esteja em branco
	assert.Equal(t, name, user.Name)   // expect user.Name seja igual à 'John Doe'.. dado mockado
	assert.Equal(t, email, user.Email) // expect user.Name seja igual à 'abc@abc.com'.. dado mockado
}

// Testar a senha
func TestUser_ValidadePassword(t *testing.T) {
	// neste teste a gente cria um usuario passando as variáveis definidas no começo.. a funcao ValidadePasword return um boolean.. então é só usar o 'assert.True()' ou 'assert.False()'
	user, err := NewUser(name, email, password)

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(password)) // Como a função ValidadePasword deverá ser aplicada a algo do tipo *User.. a gente pode usar 'user.ValidatePassword()'
	assert.False(t, user.ValidatePassword("bcde"))
	assert.NotEqual(t, password, user.Password) // esta assertion eu NAO ENTENDI O PQ..
}
