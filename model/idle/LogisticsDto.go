package idle

// LogisticsDTO 
type LogisticsDTO struct {
    // 快递/物流名称
    LogisticsName   string `json:"logistics_name,omitempty" xml:"logistics_name,omitempty"`
    // 发货人地址信息
    SenderAddress   *UserAddressDTO `json:"sender_address,omitempty" xml:"sender_address,omitempty"`
    // 物流id
    LogisticsId   string `json:"logistics_id,omitempty" xml:"logistics_id,omitempty"`
}
