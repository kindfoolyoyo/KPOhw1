package service

import "GOhw1/domain"

type Park struct {
	items   []domain.InvItem // Encapsulation: коллекция скрыта от внешнего кода.
	checker ServiceCenter    // DIP: зависимость представлена интерфейсом.
}

func NewPark(checker ServiceCenter) *Park {
	return &Park{checker: checker} // Constructor Injection: зависимость передаётся при создании Park.
}

func (p *Park) AddItem(i domain.InvItem) bool { // SRP, DRY: правила добавления собраны в одном месте.
	if i == nil || i.InvNum() == "" || !p.checker.Check(i) {
		return false
	}

	for _, item := range p.items {
		if item.InvNum() == i.InvNum() {
			return false
		}
	}

	p.items = append(p.items, i)
	return true
}

func (p *Park) TotalItems() int { // GRASP Information Expert: Park владеет списком и считает его элементы.
	return len(p.items)
}

func (p *Park) TotalEnergy() float64 { // DRY: единый алгоритм подсчёта для всех потребителей энергии.
	var total float64

	for _, item := range p.items {
		if object, ok := item.(domain.EnergySpend); ok { // Polymorphism, Low Coupling: работа через интерфейс.
			total += object.EnergySpend()
		}
	}

	return total
}

func (p *Park) EasyTransport() []domain.Transport { // GRASP Information Expert: выборка строится владельцем данных.
	var easyTransports []domain.Transport

	for _, item := range p.items {
		if transport, ok := item.(domain.Transport); ok { // OCP: алгоритм не зависит от конкретного транспорта.
			if transport.Ease() >= 6 {
				easyTransports = append(easyTransports, transport)
			}
		}
	}

	return easyTransports
}

func (p *Park) ShowAllItems() []domain.InvItem {
	showItems := make([]domain.InvItem, len(p.items))
	copy(showItems, p.items) // Encapsulation: наружу возвращается копия среза.
	return showItems
}
