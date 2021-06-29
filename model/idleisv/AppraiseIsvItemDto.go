package idleisv

// AppraiseIsvItemDTO 
type AppraiseIsvItemDTO struct {
    // 商品ID
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 商品图片,绝对途径
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    // 商品价格，单位分
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    // 商品标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
}
