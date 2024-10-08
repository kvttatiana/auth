package main

import "fmt"

func main() {
	for {
		var action int
		fmt.Println("Что вы хотите сделать?")
		fmt.Println("1. Зарегистрироваться")
		fmt.Println("2. Войти")
		fmt.Println("3. Выйти из программы")
		fmt.Scan(&action)
		if action == 1 {
			var login string
			fmt.Println("Введите логин:")
			fmt.Scan(&login)
			for {
				var password string
				fmt.Println("Введите пароль:")
				fmt.Scan(&password)

				var password1 string
				fmt.Println("Повторите пароль:")
				fmt.Scan(&password1)
				if password == password1 {
					fmt.Println("Регистрация успешно завершена")
					break
				} else {
					fmt.Println("Пароли не совпадают, повторите попытку")
				}
			}

		}
		if action == 2 {
			var login string
			fmt.Println("Введите логин:")
			fmt.Scan(&login)
			var password string
			fmt.Println("Введите пароль:")
			fmt.Scan(&password)
			fmt.Println("Вы успешно вошли в программу")

		}
		if action == 3 {
			fmt.Println("Пока пидрила")
			break
		}
	}

}
