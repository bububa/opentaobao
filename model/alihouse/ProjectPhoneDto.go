package alihouse

// ProjectPhoneDTO 
type ProjectPhoneDTO struct {
    // 分机号
    SubPhone   string `json:"sub_phone,omitempty" xml:"sub_phone,omitempty"`
    // 主号码
    MainPhone   string `json:"main_phone,omitempty" xml:"main_phone,omitempty"`
    // 1-楼盘  2-小区  3-房源  4-VR
    PhoneType   string `json:"phone_type,omitempty" xml:"phone_type,omitempty"`
    // 外部楼盘ID
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}
