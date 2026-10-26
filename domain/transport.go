package domain

type Transport interface { // ISP: содержит только контракт, необходимый транспорту.
	InvItem // Composition over Inheritance: общий контракт подключён через композицию интерфейсов.
	Ease() int
}

// Обычный велосипед
type Bike struct { // LSP: Bike может использоваться везде, где ожидается Transport.
	Num       string
	EaseScore int
	Working   bool
}

func (b Bike) InvNum() string {
	return b.Num
}

func (b Bike) Ease() int {
	return b.EaseScore
}

func (b Bike) Name() string {
	return "Велосипед"
}

func (b Bike) IsWorking() bool {
	return b.Working
}

// Электровелосипед
type EBike struct { // LSP, OCP: реализует Transport и расширяет систему без изменения Park.
	Num       string
	EaseScore int
	Energy    float64
	Working   bool
}

func (e EBike) InvNum() string {
	return e.Num
}

func (e EBike) Ease() int {
	return e.EaseScore
}

func (e EBike) EnergySpend() float64 {
	return e.Energy
}

func (e EBike) Name() string {
	return "Электровелосипед"
}

func (e EBike) IsWorking() bool {
	return e.Working
}

// Электросамокат
type EScooter struct { // LSP, OCP: реализует существующие контракты Transport и EnergySpend.
	Num       string
	EaseScore int
	Energy    float64
	Working   bool
}

func (s EScooter) InvNum() string {
	return s.Num
}

func (s EScooter) Ease() int {
	return s.EaseScore
}

func (s EScooter) EnergySpend() float64 {
	return s.Energy
}

func (s EScooter) Name() string {
	return "Электросамокат"
}

func (s EScooter) IsWorking() bool {
	return s.Working
}
