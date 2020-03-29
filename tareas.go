package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) addToList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) removeToList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printList() {

	for _, task := range t.tasks {
		fmt.Println("Name: ", task.name)
		fmt.Println("Desc: ", task.desc)
	}

}

func (t *taskList) printCompletedTaskList() {

	for _, task := range t.tasks {
		if task.completed {
			fmt.Println("Name: ", task.name)
			fmt.Println("Desc: ", task.desc)
		}
	}

}

type task struct {
	name      string
	desc      string
	completed bool
}

//* because you need update the value for the reference task.
//If you dont put *, you are going to create a copy of task, and then, you dont update de value.
func (t *task) isCompleted() {
	t.completed = true
}

func (t *task) updateDesc(newDesc string) {
	t.desc = newDesc
}

func (t *task) updateName(newName string) {
	t.name = newName
}

func main() {

	t1 := &task{ //Pointer to struct
		name: "You should be finish this course of Go",
		desc: "Desc Go",
	}

	t2 := &task{ //Pointer to struct
		name: "You should be finish this course of Python",
		desc: "Desc Pyton",
	}

	t3 := &task{ //Pointer to struct
		name: "You should be finish this course of Scala",
		desc: "Desc Scala",
	}

	listTasksNestor := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	listTasksNestor.addToList(t3)
	listTasksNestor.tasks[0].isCompleted()
	listTasksNestor.printCompletedTaskList()

	taskMaps := make(map[string]*taskList)

	taskMaps["Nestor"] = listTasksNestor

	t4 := &task{ //Pointer to struct
		name: "You should be finish this course of Java",
		desc: "Desc Java",
	}

	t5 := &task{ //Pointer to struct
		name: "You should be finish this course of Ruby",
		desc: "Desc Ruby",
	}

	listTasksPerez := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	taskMaps["Perez"] = listTasksPerez

	fmt.Println("Lista de tareas de Nestor:")
	taskMaps["Nestor"].printList()

	fmt.Println("Lista de tareas de Perez:")
	taskMaps["Perez"].printList()

}
