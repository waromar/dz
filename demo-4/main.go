package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
)
type account struct {
	login string
	password string
	url string
}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")
func main() {
	
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")
	
	account1, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN")
		return 
	}
	account1.outputPassword()

}

func promtData(promt string) (string) {
		fmt.Println(promt)
		var res string
		fmt.Scanln(&res)
		return res
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int){
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("Invalid login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	newAcc := &account{
		login: login,
		url: urlString,
		password: password,
	} 
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}