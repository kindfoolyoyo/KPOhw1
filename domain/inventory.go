package domain

type InvItem interface { // ISP: небольшой интерфейс только с общими возможностями всех объектов.
	InvNum() string
	Name() string
	IsWorking() bool
}

type EnergySpend interface { // ISP: энергопотребление вынесено в отдельный специализированный контракт.
	EnergySpend() float64
}

// Шлем
type Helmet struct { // SRP: модель хранит только данные и поведение шлема.
	Num     string
	Working bool
}

func (h Helmet) InvNum() string {
	return h.Num
}

func (h Helmet) Name() string {
	return "Шлем"
}

func (h Helmet) IsWorking() bool {
	return h.Working
}

// Зарядная станция
type EnergyStation struct { // OCP: новый потребитель энергии подключается через EnergySpend.
	Num     string
	Energy  float64
	Working bool
}

func (es EnergyStation) InvNum() string {
	return es.Num
}

func (es EnergyStation) Name() string {
	return "Зарядная станция"
}

func (es EnergyStation) EnergySpend() float64 {
	return es.Energy
}

func (es EnergyStation) IsWorking() bool {
	return es.Working
}
