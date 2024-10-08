package main

import "fmt"

func main() {

	for {
		var login string
		var password string
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

				fmt.Println("Введите пароль:")
				fmt.Scan(&password)
				// хуй
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
			var login1 string
			var password1 string
			fmt.Println("Введите логин:")
			fmt.Scan(&login1)
			if login1 == login {
				for {
					fmt.Println("Введите пароль:")
					fmt.Scan(&password)
					if password == password1 {
						fmt.Println("Вы успешно вошли в программу")

					} else {
						fmt.Println("Пароль неверный, повторите попытку")
					}
				}
			} else {
				fmt.Println("Вы не зарегистрированы")

			}

		}
		if action == 3 {
			fmt.Println("Пока пидрила")
			break
		}
	}

}
