package tmic

// ItemBo 
type ItemBo struct {
    // 选项所对应的图片cdn地址
    Img   string `json:"img,omitempty" xml:"img,omitempty"`
    // 该选项的唯一编码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 该选项的说明
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
}
