package cainiaohandover

// OpenPackageParam 结构体
type OpenPackageParam struct {
	// 包裹长度
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 包裹宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 包裹高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 包裹重量
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 包裹价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 商品参数
	ItemParams []OpenItemParam `json:"item_params,omitempty" xml:"item_params>open_item_param,omitempty"`
	// 包裹价格币种，CNY：人民币、USD：美元、RUB：卢布。
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
}
