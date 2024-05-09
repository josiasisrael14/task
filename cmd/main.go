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

func main() {
	tasks := make(map[int]Task)

	lastId := 0

	for {
		fmt.Println("ingrese las opciones")
		fmt.Println("1. ver tarea")
		fmt.Println("2. agregar tarea")
		fmt.Println("3. marca tarea  tarea completa")
		fmt.Println("4. eliminar tareas")
		fmt.Println("5. salir")
		fmt.Print("option :")

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()

		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			if len(tasks) == 0 {
				fmt.Println("no hay tareas pendientes")
			}
			for id, tarea := range tasks {
				fmt.Printf(" id :%d , tarea :%s ,description :%s completado: %t", id, tarea.Name, tarea.Description, tarea.Complet)
			}
		case 2:
			fmt.Println("nombre de la tarea")
			scanner.Scan()
			name := scanner.Text()

			fmt.Println("la descripcion de la tarea")
			scanner.Scan()
			description := scanner.Text()

			lastId++

			tasks[lastId] = Task{
				Name:        name,
				Description: description,
				Complet:     false,
			}

		case 3:
			fmt.Print("ingrese el id de la tarea completada")
			scanner.Scan()

			id, _ := strconv.Atoi(scanner.Text())

			if task, ok := tasks[id]; ok {

				if !task.Complet {
					task.Complet = true
					tasks[id] = task

					fmt.Println("tarea completada")
				} else {
					fmt.Println(" tarea  ya esta completada")
				}

			} else {
				fmt.Println("no existe la tarea")
			}

		case 4:
			fmt.Println("ingrese el id de la tarea a eliminar")
			scanner.Scan()

			id, _ := strconv.Atoi(scanner.Text())

			if task, ok := tasks[id]; ok {

				delete(tasks, id)
				fmt.Println("tarea eliminada", task)
			} else {
				fmt.Print("no existe la tarea")
			}

		case 5:
			fmt.Println("salir......")
			os.Exit(0)

		}
	}

}
