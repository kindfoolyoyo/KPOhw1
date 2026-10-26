package app

import (
	"testing"

	"GOhw1/domain"
)

func TestNewContainer(t *testing.T) {
	container := NewContainer()

	if container == nil {
		t.Fatal("Контейнер не должен быть nil")
	}

	if container.Park == nil {
		t.Fatal("Парк не должен быть nil")
	}

	scooter := domain.EScooter{
		Num:       "S001",
		EaseScore: 8,
		Energy:    0.8,
		Working:   true,
	}

	if !container.Park.AddItem(scooter) {
		t.Fatal("Исправный самокат должен приниматься")
	}

	if container.Park.TotalItems() != 1 {
		t.Errorf(
			"Ожидали 1 устройство, получили %d",
			container.Park.TotalItems(),
		)
	}
}
