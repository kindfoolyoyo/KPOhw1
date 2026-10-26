package service

import (
	"GOhw1/domain"
)

type ServiceCenter interface { // DIP: Park зависит от абстракции проверки, а не от реализации.
	Check(t domain.InvItem) bool
}

type FuncServiceCenter struct{} // SRP, GRASP Pure Fabrication: инкапсулирует правила техосмотра.

func (sc FuncServiceCenter) Check(i domain.InvItem) bool {
	if i == nil {
		return false
	}
	if !i.IsWorking() {
		return false
	}

	if t, ok := i.(domain.Transport); ok { // OCP, Polymorphism: проверка работает через Transport.
		if t.Ease() < 1 || t.Ease() > 10 {
			return false
		}
	}

	if object, ok := i.(domain.EnergySpend); ok { // OCP, Polymorphism: тип объекта перечислять не нужно.
		if object.EnergySpend() < 0 {
			return false
		}
	}

	return true
}
