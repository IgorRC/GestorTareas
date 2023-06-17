package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	complet     bool
}

type ListTask struct {
	task []Task
}

func (listTask *ListTask) addTasck(task Task) {
	listTask.task = append(listTask.task, task)
}

func (listTask *ListTask) markComplet(index int) {
	listTask.task[index].complet = true
}

func (listTask *ListTask) editTask(index int, task Task) {
	listTask.task[index] = task
}

func (listTask *ListTask) deleteTask(index int) {
	listTask.task = append(listTask.task[:index], listTask.task[index+1:]...)
}

func main() {
	listMytask := ListTask{}

	leer := bufio.NewReader(os.Stdin)
	for {
		var option int
		fmt.Println("Seleccione una opcion :\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")
		fmt.Print("Ingrese la iopcion :")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var task Task
			fmt.Print("Ingerse el nombre de la tarea :")
			task.name, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea :")
			task.description, _ = leer.ReadString('\n')
			listMytask.addTasck(task)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada :")
			fmt.Scanln(&index)
			listMytask.markComplet(index)
			fmt.Println("Tarea marcada correctamente")
		case 3:
			var task Task
			var index int
			fmt.Println("Ingrese el indice de la tarea :")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nuevo nombre :")
			task.name, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la nueva descripcion :")
			task.description, _ = leer.ReadString('\n')
			listMytask.editTask(index, task)
			fmt.Println("Tarea editada correctamente")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea :")
			fmt.Scanln(&index)
			listMytask.deleteTask(index)
			fmt.Println("Tarea eliminada correctamente")
		case 5:
			fmt.Println("Saleindo del sistema ...")
			return
		default:
			fmt.Println("Esta opcion no se encuentra disponible")
		}
		fmt.Println("LISTA DE TAREAS")
		fmt.Println("====================")
		for index, task := range listMytask.task {
			fmt.Printf("%d. %s - %s - Completado : %t\n",
				index,
				task.name,
				task.description,
				task.complet)
		}
		fmt.Println("====================")
	}
}
