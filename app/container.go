package app

import "GOhw1/service"

type Container struct { // DI Container, GRASP Pure Fabrication: отвечает за сборку зависимостей.
	Park *service.Park
}

func NewContainer() *Container { // GRASP Creator: создаёт и связывает объекты приложения.
	checker := service.FuncServiceCenter{} // DIP: конкретная реализация выбирается на уровне сборки.
	park := service.NewPark(checker)       // DI: ServiceCenter внедряется в Park через конструктор.

	return &Container{Park: park}
}
