package maitix

// PriceInfoDto 结构体
type PriceInfoDto struct {
	// 套票构成
	OpenCombinePrices []OpenCombinePriceDto `json:"open_combine_prices,omitempty" xml:"open_combine_prices>open_combine_price_dto,omitempty"`
	// 价格名称
	PriceName string `json:"price_name,omitempty" xml:"price_name,omitempty"`
	// 价格ID
	PriceId int64 `json:"price_id,omitempty" xml:"price_id,omitempty"`
	// 票品的类型 0普通单票 1固定套票  2 自由组合套票
	PriceType int64 `json:"price_type,omitempty" xml:"price_type,omitempty"`
	// 票价
	Price *Money `json:"price,omitempty" xml:"price,omitempty"`
	// 售卖最大库存数量,没有含义,以是否可售为准
	MaxStock int64 `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 场次ID
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 票品子状态0 缺货登记,1开售提醒,2 售馨,只有able_sell为false里这个才可能有值
	SubStatus int64 `json:"sub_status,omitempty" xml:"sub_status,omitempty"`
	// 是否可售 true可售 false不可售
	AbleSell bool `json:"able_sell,omitempty" xml:"able_sell,omitempty"`
}
