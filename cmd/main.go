package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Name        string
	Description string
	Complet     bool
}

type Usuario struct {
	IdUser     int
	NameUser   string
	Password   string
	Priveligio string
}

func main() {
	tasks := make(map[int]Task)
	usuario := make(map[int]Usuario)
	lastId := 0
	idUser := 0

	var (
		nameUser          string
		password          string
		privilegio        string
		privilegioUsuario string
	)

	// Agregar un usuario administrador inicial
	usuario[0] = Usuario{
		IdUser:     0,
		NameUser:   "jairo",
		Password:   "pamela",
		Priveligio: "administrador",
	}

	for {
		encontrado := false

		fmt.Println("Ingrese Usuario:")
		fmt.Scanln(&nameUser)

		fmt.Println("Ingrese Password:")
		fmt.Scanln(&password)

		for _, u := range usuario {
			if u.NameUser == nameUser && u.Password == password {
				encontrado = true
				privilegioUsuario = u.Priveligio
				break
			}
		}

		if !encontrado {
			fmt.Println("Usuario no existe o contraseña incorrecta")
			continue
		}

		if privilegioUsuario == "administrador" {
			for {
				fmt.Println("Bienvenido administrador")
				fmt.Println("Ingrese las opciones:")
				fmt.Println("1. Ver tarea")
				fmt.Println("2. Agregar tarea")
				fmt.Println("3. Marcar tarea como completa")
				fmt.Println("4. Eliminar tarea")
				fmt.Println("5. Agregar usuario")
				fmt.Println("6. Ver usuarios")
				fmt.Println("7. Salir")
				fmt.Print("Opción: ")

				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				choice, _ := strconv.Atoi(scanner.Text())

				switch choice {
				case 1:
					if len(tasks) == 0 {
						fmt.Println("No hay tareas pendientes")
					} else {
						for id, tarea := range tasks {
							fmt.Printf("ID: %d, Tarea: %s, Descripción: %s, Completado: %t\n", id, tarea.Name, tarea.Description, tarea.Complet)
						}
					}
				case 2:
					fmt.Println("Nombre de la tarea:")
					scanner.Scan()
					name := scanner.Text()

					fmt.Println("Descripción de la tarea:")
					scanner.Scan()
					description := scanner.Text()

					lastId++
					tasks[lastId] = Task{
						Name:        name,
						Description: description,
						Complet:     false,
					}

				case 3:
					fmt.Println("Ingrese el ID de la tarea completada:")
					scanner.Scan()
					id, _ := strconv.Atoi(scanner.Text())

					if task, ok := tasks[id]; ok {
						if !task.Complet {
							task.Complet = true
							tasks[id] = task
							fmt.Println("Tarea completada")
						} else {
							fmt.Println("Tarea ya está completada")
						}
					} else {
						fmt.Println("No existe la tarea")
					}

				case 4:
					fmt.Println("Ingrese el ID de la tarea a eliminar:")
					scanner.Scan()
					id, _ := strconv.Atoi(scanner.Text())

					if task, ok := tasks[id]; ok {
						delete(tasks, id)
						fmt.Println("Tarea eliminada:", task)
					} else {
						fmt.Println("No existe la tarea")
					}

				case 5:
					fmt.Println("Ingrese el usuario a agregar:")
					fmt.Scanln(&nameUser)

					fmt.Println("Ingrese el password:")
					fmt.Scanln(&password)

					fmt.Println("Ingrese el privilegio:")
					fmt.Scanln(&privilegio)

					idUser++
					usuario[idUser] = Usuario{
						IdUser:     idUser,
						NameUser:   nameUser,
						Password:   password,
						Priveligio: privilegio,
					}

					fmt.Println("Usuario agregado:", usuario[idUser])

				case 6:
					if len(usuario) == 0 {
						fmt.Println("No hay usuarios agregados")
					} else {
						for id, usuarios := range usuario {
							fmt.Printf("ID: %d, Usuario: %s, Password: %s, Privilegio: %s\n", id, usuarios.NameUser, usuarios.Password, usuarios.Priveligio)
						}
					}

				case 7:
					fmt.Println("Salir...")
					os.Exit(0)
				}
			}
		} else if privilegioUsuario == "usuario" {
			for {
				fmt.Println("Ingrese las opciones:")
				fmt.Println("1. Ver tarea")
				fmt.Println("2. Marcar tarea como completa")
				fmt.Println("3. Salir")
				fmt.Print("Opción: ")

				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				choice, _ := strconv.Atoi(scanner.Text())

				switch choice {
				case 1:
					if len(tasks) == 0 {
						fmt.Println("No hay tareas pendientes")
					} else {
						for id, tarea := range tasks {
							fmt.Printf("ID: %d, Tarea: %s, Descripción: %s, Completado: %t\n", id, tarea.Name, tarea.Description, tarea.Complet)
						}
					}
				case 2:
					fmt.Println("Ingrese el ID de la tarea completada:")
					scanner.Scan()
					id, _ := strconv.Atoi(scanner.Text())

					if task, ok := tasks[id]; ok {
						if !task.Complet {
							task.Complet = true
							tasks[id] = task
							fmt.Println("Tarea completada")
						} else {
							fmt.Println("Tarea ya está completada")
						}
					} else {
						fmt.Println("No existe la tarea")
					}

				case 3:
					fmt.Println("Salir...")
					os.Exit(0)
				}
			}
		} else {
			fmt.Println("Privilegio no reconocido")
		}
	}
}
