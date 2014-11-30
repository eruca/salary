package doc

import (
	"encoding/json"
	"io/ioutil"
)

var (
	DocCnt     = 11
	configfile = "./doc.json"
	dotFees    = 0.0
	outFees    = 0.0
	nightFees  = 0.0
	fineFees   = 0.0
	perDoc     = 0.0
	Docs       = []Doc{}
)

type Doc struct {
	Name       string  `json:"name"`
	Pinyin     string  `json:"pinyin"`
	Point      float32 `json:"point"`
	on         bool
	attendRate float32

	nights    int //夜班数
	patients  int //病人数
	delayDays int
	delayCnt  int
	salary    float32
}

func NewDoc(name string, pt float32) *Doc {
	return &Doc{Name: name, Point: pt, on: true, attendRate: 1.0}
}

func LoadConfig() []Doc {
	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &Docs)
	if err != nil {
		panic(err)
	}
	return Docs
}

func GetBonus(bonus, outfee, nightfee, perdoc float32) {
	dotFees = bonus - outfee - nightfee - perdoc*DocCnt
	outFees = outfee
	nightFees = nightfee
	perDoc = perdoc
	if bonus != (dotFees + outfee + nightFees + perDoc*DocCnt) {
		panic("总奖金数目不对")
	}
}
