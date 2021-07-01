package game

// AppleCardDto 结构体
type AppleCardDto struct {
	// 面值
	FacePrice string `json:"face_price,omitempty" xml:"face_price,omitempty"`
	// 有效期
	Expire string `json:"expire,omitempty" xml:"expire,omitempty"`
	// 卡密
	CardPass string `json:"card_pass,omitempty" xml:"card_pass,omitempty"`
	// 卡号
	CardNo string `json:"card_no,omitempty" xml:"card_no,omitempty"`
	// 产品编码
	ZhxGoodsId string `json:"zhx_goods_id,omitempty" xml:"zhx_goods_id,omitempty"`
}
