package ascpchannel

// ExternalSenderRequest 
type ExternalSenderRequest struct {
    // 发货人 手机号
    MobilePhone   string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
    // 发货人名称
    ContactName   string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
}
