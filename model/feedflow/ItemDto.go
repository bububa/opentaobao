package feedflow

// ItemDto 
type ItemDto struct {
    // 商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 商品标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 商品落地页
    LinkUrl   string `json:"link_url,omitempty" xml:"link_url,omitempty"`
    // 商品主图
    ImgUrl   string `json:"img_url,omitempty" xml:"img_url,omitempty"`
    // 是否可以使用，false不可以进行广告投放
    IsAccessAllowed   bool `json:"is_access_allowed,omitempty" xml:"is_access_allowed,omitempty"`
    // 不可以使用的原因
    AccessAllowedInfo   *AccessAllowedDto `json:"access_allowed_info,omitempty" xml:"access_allowed_info,omitempty"`
}
