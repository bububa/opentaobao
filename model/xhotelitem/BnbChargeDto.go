package xhotelitem

// BnbChargeDto 结构体
type BnbChargeDto struct {
	// 是否允许加人 0不允许 1允许
	AddPeople int64 `json:"add_people,omitempty" xml:"add_people,omitempty"`
	// 允许加人数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 加人费用金额
	Fee int64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// 最小收费年龄
	MinChargingAge int64 `json:"min_charging_age,omitempty" xml:"min_charging_age,omitempty"`
}
