package main

import (
	"bufio"
	"fmt"
	"gestortareas/modelo"
	"os"
)

func main() {
	listMytask := modelo.ListTask{}

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
			var task modelo.Task
			fmt.Print("Ingerse el nombre de la tarea :")
			task.Name, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea :")
			task.Description, _ = leer.ReadString('\n')
			listMytask.AddTasck(task)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada :")
			fmt.Scanln(&index)
			listMytask.MarkComplet(index)
			fmt.Println("Tarea marcada correctamente")
		case 3:
			var task modelo.Task
			var index int
			fmt.Println("Ingrese el indice de la tarea :")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nuevo nombre :")
			task.Name, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la nueva descripcion :")
			task.Description, _ = leer.ReadString('\n')
			listMytask.EditTask(index, task)
			fmt.Println("Tarea editada correctamente")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea :")
			fmt.Scanln(&index)
			listMytask.DeleteTask(index)
			fmt.Println("Tarea eliminada correctamente")
		case 5:
			fmt.Println("Saleindo del sistema ...")
			return
		default:
			fmt.Println("Esta opcion no se encuentra disponible")
		}
		fmt.Println("LISTA DE TAREAS")
		fmt.Println("====================")
		for index, task := range listMytask.GetArrTask() {
			fmt.Printf("%d. %s - %s - Completado : %t\n",
				index,
				task.Name,
				task.Description,
				task.Complet)
		}
		fmt.Println("====================")
	}
}
