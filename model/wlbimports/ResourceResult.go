package wlbimports

// ResourceResult 
type ResourceResult struct {

    // 服务报价。首重（1磅) CNY￥3元 续重（1磅) CNY￥6元
    DeliveryPrice   *DeliveryPrice `json:"delivery_price,omitempty"`

    // 时效，单位（小时）
    DeliveryTime   int64 `json:"delivery_time,omitempty"`

    // 资源代码
    ResCode   string `json:"res_code,omitempty"`

    // 资源ID
    ResId   int64 `json:"res_id,omitempty"`

    // 魔方天下美国线
    ResourceName   string `json:"resource_name,omitempty"`

}
