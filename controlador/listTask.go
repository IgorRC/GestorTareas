package controlador

import "gestortareas/modelo"

type ListTask struct {
	task []modelo.Task
}

func (listTask *ListTask) AddTasck(task *modelo.Task) {
	listTask.task = append(listTask.task, *task)
}

func (listTask *ListTask) MarkComplet(index int) {
	listTask.task[index].Complet = true
}

func (listTask *ListTask) EditTask(index int, task *modelo.Task) {
	listTask.task[index] = *task
}

func (listTask *ListTask) DeleteTask(index int) {
	listTask.task = append(listTask.task[:index], listTask.task[index+1:]...)
}

func (listTask *ListTask) GetArrTask() []modelo.Task {
	return listTask.task
}
