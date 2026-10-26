package service

import (
	"GOhw1/domain"
	"math"
	"testing"
)

type stubServiceCenter struct { // DIP, Test Double: заменяет реальный ServiceCenter в юнит-тестах.
	approved bool
	calls    int
}

func (s *stubServiceCenter) Check(i domain.InvItem) bool {
	s.calls++
	return s.approved
}

func TestParkRejectsWhenCheckFails(t *testing.T) {
	checker := &stubServiceCenter{
		approved: false,
	}

	park := NewPark(checker) // DI: тестовая зависимость внедряется через конструктор.

	item := domain.EScooter{
		Num:       "ES001",
		EaseScore: 5,
		Energy:    0.5,
		Working:   true,
	}

	result := park.AddItem(item)

	if checker.calls != 1 {
		t.Errorf("ожидали 1 вызов техосмотра, получили %d", checker.calls)
	}

	if result {
		t.Error("парк принял устройство, которое не прошло техосмотр")
	}

	if park.TotalItems() != 0 {
		t.Error("в парке не должно быть устройств")
	}
}

func TestParkAcceptsApprovedItem(t *testing.T) {
	checker := &stubServiceCenter{
		approved: true,
	}

	park := NewPark(checker)

	scooter := domain.EScooter{
		Num:       "ES001",
		EaseScore: 5,
		Energy:    0.5,
		Working:   true,
	}

	result := park.AddItem(scooter)

	if !result {
		t.Error("парк должен принять одобренный самокат")
	}

	if park.TotalItems() != 1 {
		t.Errorf("ожидали 1 устройство, получили %d", park.TotalItems())
	}
}

func TestParkRejectsDuplicateNumber(t *testing.T) {
	checker := &stubServiceCenter{approved: true}

	park := NewPark(checker)

	first := domain.Bike{
		Num:       "B001",
		EaseScore: 5,
		Working:   true,
	}

	second := domain.EScooter{
		Num:       "B001",
		EaseScore: 5,
		Energy:    5,
		Working:   true,
	}

	park.AddItem(first)

	result := park.AddItem(second)

	if result {
		t.Error("нельзя добавить устройство с повторяющимся номером")
	}

	if park.TotalItems() != 1 {
		t.Errorf("ожидали 1 устройство, получили %d", park.TotalItems())
	}
}

func TestParkTotalEnergy(t *testing.T) {
	checker := &stubServiceCenter{approved: true}
	park := NewPark(checker)

	scooter := domain.EScooter{
		Num:       "S001",
		EaseScore: 8,
		Energy:    0.8,
		Working:   true,
	}

	station := domain.EnergyStation{
		Num:     "C001",
		Energy:  2.5,
		Working: true,
	}

	helmet := domain.Helmet{
		Num:     "H001",
		Working: true,
	}

	if !park.AddItem(scooter) ||
		!park.AddItem(station) ||
		!park.AddItem(helmet) {
		t.Fatal("Не получилось добавить предметы в парк")
	}

	true_reesult := 3.3
	result := park.TotalEnergy()

	if math.Abs(result-true_reesult) > 1e-9 {
		t.Errorf("Ожидали %f, получили %f", true_reesult, result)
	}
}

func TestParkEasyTransport(t *testing.T) {
	checker := &stubServiceCenter{approved: true}
	park := NewPark(checker)

	bike := domain.Bike{
		Num:       "B001",
		EaseScore: 5,
		Working:   true,
	}

	escooter := domain.EScooter{
		Num:       "ES001",
		EaseScore: 6,
		Energy:    0.8,
		Working:   true,
	}

	if !park.AddItem(bike) || !park.AddItem(escooter) {
		t.Fatal("Не удалось добавить транспорт")
	}

	result := park.EasyTransport()

	if len(result) != 1 {
		t.Fatalf("Ожидали 2 устройства, получили %d", len(result))
	}

	if result[0].InvNum() != "ES001" {
		t.Error("Неверный список транспорта для новичков")
	}
}

func TestParkShowAllItems(t *testing.T) {
	checker := &stubServiceCenter{approved: true}
	park := NewPark(checker)

	helmet := domain.Helmet{
		Num:     "H001",
		Working: true,
	}

	bike := domain.Bike{
		Num:       "B001",
		EaseScore: 7,
		Working:   true,
	}

	if !park.AddItem(helmet) || !park.AddItem(bike) {
		t.Fatal("Не удалось подготовить парк")
	}

	items := park.ShowAllItems()

	if len(items) != 2 {
		t.Fatalf("Ожидали 2 объекта, получили %d", len(items))
	}

	if items[0].InvNum() != "H001" || items[1].InvNum() != "B001" {
		t.Error("Инвентарные номера не совпадают")
	}
}

func TestParkEmpty(t *testing.T) {
	checker := &stubServiceCenter{approved: true}
	park := NewPark(checker)

	if park.TotalItems() != 0 {
		t.Error("Новый парк должен быть пустым")
	}

	if park.TotalEnergy() != 0 {
		t.Error("Пустой парк не должен потреблять энергию")
	}

	if len(park.EasyTransport()) != 0 {
		t.Error("В пустом парке нет транспорта для новичков")
	}

	if len(park.ShowAllItems()) != 0 {
		t.Error("В пустом парке нет инвентаря")
	}
}
