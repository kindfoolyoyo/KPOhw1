package service

import (
	"testing"

	"GOhw1/domain"
)

func TestCheckNil(t *testing.T) {
	checker := FuncServiceCenter{}

	if checker.Check(nil) {
		t.Error("nil не должен проходить техосмотр")
	}
}

func TestCheckWorkingItem(t *testing.T) {
	checker := FuncServiceCenter{}

	helmet := domain.Helmet{
		Num:     "H-002",
		Working: true,
	}

	if !checker.Check(helmet) {
		t.Error("исправный шлем должен проходить техосмотр")
	}
}
func TestChecNotWorkingItem(t *testing.T) {
	checker := FuncServiceCenter{}

	helmet := domain.Helmet{
		Num:     "H-001",
		Working: false,
	}

	if checker.Check(helmet) {
		t.Error("неисправный шлем не должен проходить техосмотр")
	}
}

func TestCheckEase(t *testing.T) {
	checker := FuncServiceCenter{}

	tests := []struct { // DRY: несколько сценариев объединены в табличный тест.
		name     string
		ease     int
		expected bool
	}{{"ниже минимума", 0, false}, {"минимум", 1, true}, {"середина", 5, true}, {"максимум", 10, true}, {"выше максимума", 11, false}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scooter := domain.EScooter{
				Num:       "ES001",
				EaseScore: tt.ease,
				Energy:    0.8,
				Working:   true,
			}

			result := checker.Check(scooter)

			if result != tt.expected {
				t.Errorf("EaseScore=%d: ожидали %t, получили %t", tt.ease, tt.expected, result)
			}
		})
	}
}

func TestCheckEnergy(t *testing.T) {
	checker := FuncServiceCenter{}

	tests := []struct { // DRY: одна тестовая логика используется для разных значений энергии.
		name     string
		energy   float64
		expected bool
	}{{"Отрицательная", -1, false}, {"Нулевая", 0, true}, {"Положительная", 2.5, true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			station := domain.EnergyStation{
				Num:     "ES001",
				Energy:  tt.energy,
				Working: true,
			}

			result := checker.Check(station)

			if result != tt.expected {
				t.Errorf("Energy=%f: ожидали %t, получили %t", tt.energy, tt.expected, result)
			}
		})
	}
}
