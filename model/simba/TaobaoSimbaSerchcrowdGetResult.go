package simba

// TaobaoSimbaSerchcrowdGetResult 
type TaobaoSimbaSerchcrowdGetResult struct {
    // 出价方式1:溢价;0:出价
    PriceMode   int64 `json:"price_mode,omitempty" xml:"price_mode,omitempty"`
    // 返回的溢价比例,乘的关系,discount=121,代表溢价21%
    Discount   int64 `json:"discount,omitempty" xml:"discount,omitempty"`
    // 人群信息
    Crowd   *CrowdDto `json:"crowd,omitempty" xml:"crowd,omitempty"`
    // 定向id,出价删除 改价,修改状态时用
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 人群是否溢价1:不溢价;0:溢价
    IsDefaultPrice   int64 `json:"is_default_price,omitempty" xml:"is_default_price,omitempty"`
    // 人群上下线状态0:暂停, 1:启用
    OnlineStatus   int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
}
