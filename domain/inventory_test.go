package domain

import "testing"

func TestItemProperties(t *testing.T) {
	tests := []struct {
		name         string
		item         InvItem
		expectedName string
		working      bool
	}{
		{"велосипед", Bike{Num: "B001", Working: true}, "Велосипед", true},
		{"электровелосипед", EBike{Num: "EB001", Working: false}, "Электровелосипед", false},
		{"самокат", EScooter{Num: "ES001", Working: true}, "Электросамокат", true},
		{"шлем", Helmet{Num: "H001", Working: false}, "Шлем", false},
		{"станция", EnergyStation{Num: "C001", Working: true}, "Зарядная станция", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.item.Name(); res != tt.expectedName {
				t.Errorf("Name(): ожидали %q, получили %q", tt.expectedName, res)
			}
			if res := tt.item.IsWorking(); res != tt.working {
				t.Errorf("IsWorking(): ожидали %t, получили %t", tt.working, res)
			}
		})
	}
}
