package handlers

import "errors"

func PanicIfUserError(err error) {
	if err != nil {
		err := errors.New("userService ----" + err.Error())
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err := errors.New("taskService ----" + err.Error())
		panic(err)
	}
}
