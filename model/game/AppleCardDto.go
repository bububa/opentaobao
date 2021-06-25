package game

// AppleCardDto 
type AppleCardDto struct {

    // 面值
    FacePrice   string `json:"face_price,omitempty"`

    // 有效期
    Expire   string `json:"expire,omitempty"`

    // 卡密
    CardPass   string `json:"card_pass,omitempty"`

    // 卡号
    CardNo   string `json:"card_no,omitempty"`

    // 产品编码
    ZhxGoodsId   string `json:"zhx_goods_id,omitempty"`

}
