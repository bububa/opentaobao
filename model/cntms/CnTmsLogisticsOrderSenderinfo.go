package cntms

// CnTmsLogisticsOrderSenderinfo 
type CnTmsLogisticsOrderSenderinfo struct {
    // 发件人省份
    SenderProvince   string `json:"sender_province,omitempty" xml:"sender_province,omitempty"`
    // 发件人城市
    SenderCity   string `json:"sender_city,omitempty" xml:"sender_city,omitempty"`
    // 发件人区县
    SenderArea   string `json:"sender_area,omitempty" xml:"sender_area,omitempty"`
    // 发件人地址
    SenderAddress   string `json:"sender_address,omitempty" xml:"sender_address,omitempty"`
    // 发件人邮编
    SenderZipCode   string `json:"sender_zip_code,omitempty" xml:"sender_zip_code,omitempty"`
    // 发件人姓名
    SenderName   string `json:"sender_name,omitempty" xml:"sender_name,omitempty"`
    // 发件人手机，手机与电话必须有一值不为空
    SenderMobile   string `json:"sender_mobile,omitempty" xml:"sender_mobile,omitempty"`
    // 发件人电话，手机与电话必须有一值不为空
    SenderPhone   string `json:"sender_phone,omitempty" xml:"sender_phone,omitempty"`
}
