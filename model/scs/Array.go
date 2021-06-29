package scs

// Array 
type Array struct {
    // 组内数据
    Datas   []Array `json:"datas,omitempty" xml:"datas>array,omitempty"`
    // 组号(调用方暂不用关心)
    GroupId   string `json:"group_id,omitempty" xml:"group_id,omitempty"`
    // 图片URL
    Image   string `json:"image,omitempty" xml:"image,omitempty"`
    // 商品ID(调用方暂不用关心)
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 创意ID(调用方暂不用关心)
    CrtId   string `json:"crt_id,omitempty" xml:"crt_id,omitempty"`
    // 模板ID(调用方暂不用关心)
    TemplateId   int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
    // 标题(调用方暂不用关心)
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 扩展信息(即抠图结果，fg字段即BASE64编码的图片二进制数据)
    ExtInfo   string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
}
