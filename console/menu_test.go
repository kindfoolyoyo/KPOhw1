package console

import (
	"GOhw1/service"
	"bytes"
	"strings"
	"testing"
)

func runMenu(input string) (*service.Park, string) {
	park := service.NewPark(service.FuncServiceCenter{})
	var output bytes.Buffer
	Run(park, strings.NewReader(input), &output)
	return park, output.String()
}

func TestMenuAddScooter(t *testing.T) {
	input := "1\n 3\n ES001\n 1\n8 \n0.8 \n0 \n"
	park, output := runMenu(input)

	if park.TotalItems() != 1 {
		t.Errorf("Ожидали 1 предмет, получили %d", park.TotalItems())
	}

	if !strings.Contains(output, "Предмет принят в парк") {
		t.Error("Нет сообщения об успешной приёмке")
	}

	items := park.ShowAllItems()
	if len(items) != 1 {
		t.Fatalf("Ожидали 1 предмет, получили %d", len(items))
	}

	if items[0].InvNum() != "ES001" || items[0].Name() != "Электросамокат" {
		t.Error("В парк добавлен не тот предмет")
	}
}

func TestMenuRejectBrokenHelmet(t *testing.T) {
	input := "1 \n4 \nH001 \n0 \n0 \n"
	park, output := runMenu(input)

	if park.TotalItems() != 0 {
		t.Error("Неисправный шлем не должен попадать в парк")
	}

	if !strings.Contains(output, "Предмет не прошёл приёмку") {
		t.Error("Нет сообщения об отказе в приёмке")
	}
}

func TestMenuShowInformation(t *testing.T) {
	input := "1 \n3 \nES001 \n1 \n8 \n2 \n2 \n3 \n4 \n5 \n0 \n"
	park, output := runMenu(input)

	if park.TotalItems() != 1 {
		t.Fatalf("Ожидали 1 предмет, получили %d", park.TotalItems())
	}

	if !strings.Contains(output, "Количество предметов в парке -  1") {
		t.Error("Количество предметов отображается неверно")
	}

	if !strings.Contains(output, "Суммарное суточное энергопотребление -  2") {
		t.Error("Энергопотребление отображается неверно")
	}

	if strings.Count(output, "Электросамокат ES001") != 2 {
		t.Error("Самокат должен отображаться в списке для новичков и в общем списке")
	}
}

func TestMenuInvalidItemType(t *testing.T) {
	input := "1 \n9 \n0 \n"
	park, output := runMenu(input)

	if park.TotalItems() != 0 {
		t.Error("Предмет с неверным типом не должен добавляться")
	}

	if !strings.Contains(output, "Неверный ввод") {
		t.Error("Нет сообщения о неверном вводе")
	}
}
