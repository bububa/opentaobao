package tuanhotel

// TuanItemRelateGiftVO 结构体
type TuanItemRelateGiftVO struct {
	// 数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 类别：1-温泉；2-滑雪；3-门票；4-美食；5-下午茶；6-SPA；7-礼包；8-亲子；9-交通；10-代金券；999-用户自定义
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 宝贝ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
}
