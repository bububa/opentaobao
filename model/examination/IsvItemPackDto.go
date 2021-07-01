package examination

// IsvItemPackDto 
type IsvItemPackDto struct {
    // 加项包id
    IsvPackId   string `json:"isv_pack_id,omitempty" xml:"isv_pack_id,omitempty"`
    // 加项包名称
    PackName   string `json:"pack_name,omitempty" xml:"pack_name,omitempty"`
    // 售价，单位分
    SoldPrice   string `json:"sold_price,omitempty" xml:"sold_price,omitempty"`
    // 类型：1、可选加项包 2、声明式加项包.
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 加项包包含的单项id列表
    IsvItemIds   []string `json:"isv_item_ids,omitempty" xml:"isv_item_ids>string,omitempty"`
    // 版本号，防止isv更改未同步给健康，提供给isv做校验的
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    // 结算价，单位分
    SettlePrice   string `json:"settle_price,omitempty" xml:"settle_price,omitempty"`
}
