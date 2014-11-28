package doctor

const (
	doctors     = 11
	perNightFee = 20.0
	perPerson   = 400.0
)

type Doctor struct {
	Name     string
	Point    float32
	Nights   int64   //夜班数
	Patients int64   //病人数
	Fine     float32 //病例迟交罚钱
	On       bool
	Salary   float32
}

func NewDoctor(name string, pt float32) *Doctor {
	return &Doctor{Name: name, Point: pt, On: true}
}

func (this *Doctor) TurnOn() {
	this.On = true
}

func (this *Doctor) TurnOff() {
	this.On = false
}
