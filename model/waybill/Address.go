package waybill

// Address 
type Address struct {

    // 市
    City   string `json:"city,omitempty"`

    // 详细地址
    Detail   string `json:"detail,omitempty"`

    // 区
    District   string `json:"district,omitempty"`

    // 省
    Province   string `json:"province,omitempty"`

    // 街道
    Town   string `json:"town,omitempty"`

}
