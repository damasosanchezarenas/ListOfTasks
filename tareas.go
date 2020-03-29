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
		desc: "Its necesary completed this course for go to the next level.",
	}

	t2 := &task{ //Pointer to struct
		name: "You should be finish this course of Python",
		desc: "Its necesary completed this course for go to the next level.",
	}

	t3 := &task{ //Pointer to struct
		name: "You should be finish this course of Scala",
		desc: "Its necesary completed this course for go to the next level.",
	}

	listTasks := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	listTasks.addToList(t3)
	listTasks.printList()
	listTasks.tasks[0].isCompleted()
	println("TAREAS COMPLETADAS: ")
	listTasks.printCompletedTaskList()

}
