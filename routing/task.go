package routing

import (
	. "../controllers"
	. "../repositories"
)

var (
	taskRoutes Routes

	controller *TaskController
	endPoint = "/tasks"
)

func init() {

	controller = &TaskController{ Repo: &TasksRepository{} }
	controller.Repo.Connect()

	taskRoutes = Routes {

		Route {
			"Hello",
			"GET",
			"/hello",
			controller.HelloWorld,
		},
		Route {
			"GetAllTasks",
			"GET",
			endPoint,
			controller.FindAllTasks,
		},
		Route {
			"GetTaskById",
			"GET",
			endPoint + "/{id}",
			controller.FindTaskById,
		},
		Route {
			"CreateTask",
			"POST",
			endPoint,
			controller.CreateTask,
		},
		Route {
			"UpdateTask",
			"UPDATE",
			endPoint,
			controller.UpdateTask,
		},
		Route {
			"DeleteTask",
			"DELETE",
			endPoint + "/{id}",
			controller.DeleteTask,
		},
	}
}