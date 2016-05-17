package model

import ()
import (
	"fmt"
	"log"
	"errors"
)

const (
	DB_TABLE = "USER"
	DB_CL_Name = "Name"
	DB_CL_Email = "Email"
	DB_CL_Password = "Password"
	DB_IDS = DB_TABLE + ":" + "IDS"
	DB_KEY = DB_TABLE + ":" + DB_CL_Name
)

type User struct {
	Id       uint64
	Name     string
	Email    string
	Password string
	Client
}

func (u *User) GetId() uint64 {
	return u.Id
}
func (u *User) GetName() string {
	return u.Name
}
func (u *User) SetName(name string) {
	u.Name = name
}
func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) SetEmail(mail string) {
	u.Email = mail
}
func (u *User) GetPassWord() string {
	return u.Password
}
func (u *User) SetPassword(pass string) {
	u.Password = pass
}
func (u *User) OpenConnect() {
	u.Client.Init()
}
func (u *User) Println() {
	log.Println("uesr:", u.GetName(), "- pass:", u.GetPassWord())
}
func (u *User) CheckExistByName() bool {

	var key = u.GetKey()
	var result map[string]string
	result, _ = u.GetClient().HGetAll(key)

	if len(result) > 0 {
		return true
	} else {
		return false
	}
}
func (u *User)GetKey() string {
	var key = DB_TABLE + ":" + u.GetName()
	return key
}
func (u *User) Add() error {

	u.OpenConnect()

	var err error
	var key = u.GetKey()

	// add new record
	if u.CheckExistByName() {

		err = errors.New("Existed User")
	} else {

		var iResult int64
		// 1. add primary key
		iResult, _ = u.GetClient().SAdd(DB_IDS, u.GetName())

		// 2. add detail record
		if iResult == 0 {
			fmt.Println("exist:", u.GetName())
		} else {
			pairs := make(map[string]string)
			pairs[DB_CL_Name] = u.GetName()
			pairs[DB_CL_Password] = u.GetPassWord()
			pairs[DB_CL_Email] = u.GetEmail()
			_ = u.GetClient().HMSet(key, pairs)
		}
	}

	return err
}

