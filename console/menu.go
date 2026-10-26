package console

import (
	"GOhw1/domain"
	"GOhw1/service"
	"fmt"
	"io"
)

func Run(park *service.Park, input io.Reader, output io.Writer) { // SRP, GRASP Controller: принимает команды и вызывает операции Park.
	for {
		fmt.Fprintln(output, "Добро пожаловать в ВышКат, друг!")
		fmt.Fprintln(output, "0 - выход")
		fmt.Fprintln(output, "1 - добавить транспорт")
		fmt.Fprintln(output, "2 - количество обектов")
		fmt.Fprintln(output, "3 - потребление энергии")
		fmt.Fprintln(output, "4 - показать транспорт для новичков")
		fmt.Fprintln(output, "5 - показать все объекты")

		var command int
		fmt.Fprint(output, "Выберите действие ")
		if _, err := fmt.Fscan(input, &command); err != nil {
			fmt.Fprintln(output, "Неверный ввод")
			return
		}

		switch command {
		case 0:
			return
		case 1:
			addItem(park, input, output)
		case 2:
			fmt.Fprintln(output, "Количество предметов в парке - ", park.TotalItems())
		case 3:
			fmt.Fprintln(output, "Суммарное суточное энергопотребление - ", park.TotalEnergy())
		case 4:
			if len(park.EasyTransport()) == 0 {
				fmt.Fprintln(output, "Подходящего транспорта нет")
			}

			for _, t := range park.EasyTransport() {
				fmt.Fprintln(output, t.Name(), t.InvNum())
			}

		case 5:
			if park.TotalItems() == 0 {
				fmt.Fprintln(output, "Парк пуст")
			}

			for _, item := range park.ShowAllItems() {
				fmt.Fprintln(output, item.Name(), item.InvNum())
			}

		default:
			return
		}
	}
}

func addItem(park *service.Park, input io.Reader, output io.Writer) { // GRASP Creator: создаёт объект из введённых пользователем данных.
	var kind int

	fmt.Fprintln(output, "Выберите предмет")
	fmt.Fprintln(output, "1 - велосипед")
	fmt.Fprintln(output, "2 - электровелосипед")
	fmt.Fprintln(output, "3 - самокат")
	fmt.Fprintln(output, "4 - шлем")
	fmt.Fprintln(output, "5 - зарядная станция")

	if _, err := fmt.Fscan(input, &kind); err != nil {
		fmt.Fprintln(output, "Ошибка считывания")
		return
	}

	if kind <= 0 || kind >= 6 {
		fmt.Fprintln(output, "Неверный ввод")
		return
	}

	fmt.Fprintln(output, "Выбран предмет:", kind)

	var num string
	var intworking int

	fmt.Fprintln(output, "Введите нвентарный номер ")
	if _, err := fmt.Fscan(input, &num); err != nil {
		fmt.Fprintln(output, "Ошибка считывания")
		return
	}

	fmt.Fprintln(output, "Введите состояние исправности: 1 - исправен, 0 - неисправен ")
	if _, err := fmt.Fscan(input, &intworking); err != nil {
		fmt.Fprintln(output, "Ошибка считывания")
		return
	}

	if intworking != 0 && intworking != 1 {
		fmt.Fprintln(output, "Неверный ввод")
		return
	}

	working := intworking == 1

	var energy float64
	var ease int

	if kind == 1 || kind == 2 || kind == 3 {
		fmt.Fprintln(output, "Введите уровень простоты управления (от 1 до 10) ")
		if _, err := fmt.Fscan(input, &ease); err != nil {
			fmt.Fprintln(output, "Ошибка считывания")
			return
		}
	}

	if kind == 2 || kind == 3 || kind == 5 {
		fmt.Fprint(output, "Введите суточное энергопотребление: ")

		if _, err := fmt.Fscan(input, &energy); err != nil {
			fmt.Fprintln(output, "Ошибка считывания")
			return
		}
	}

	var item domain.InvItem // LSP: любой объект парка используется через общий контракт InvItem.

	switch kind {
	case 1:
		item = domain.Bike{
			Num:       num,
			EaseScore: ease,
			Working:   working,
		}
	case 2:
		item = domain.EBike{
			Num:       num,
			EaseScore: ease,
			Working:   working,
			Energy:    energy,
		}
	case 3:
		item = domain.EScooter{
			Num:       num,
			EaseScore: ease,
			Working:   working,
			Energy:    energy,
		}
	case 4:
		item = domain.Helmet{
			Num:     num,
			Working: working,
		}
	case 5:
		item = domain.EnergyStation{
			Num:     num,
			Working: working,
			Energy:  energy,
		}

	}

	if park.AddItem(item) { // SoC: правила приёмки находятся в Park, а не в консольном слое.
		fmt.Fprintln(output, "Предмет принят в парк")
	} else {
		fmt.Fprintln(output, "Предмет не прошёл приёмку")
	}
}
