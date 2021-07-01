package btrip

// PassengerQuantityRq 结构体
type PassengerQuantityRq struct {
	// 乘客类型;ADT:"普通成人", CHD:"儿童", STU:"留学生", LABOR:"劳工", MIGRANT:"新移民", MARINER:"海员", OLD:"老人", YOUNG:"青年", INFANT:"婴儿", OTHER:"特殊身份"
	PassengerType string `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 人员数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
