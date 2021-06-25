package cntms

// CnTmsLogisticsOrderReceiverInfo 
type CnTmsLogisticsOrderReceiverInfo struct {

    // 收件人省份
    ReceiverProvince   string `json:"receiver_province,omitempty"`

    // 收件人城市
    ReceiverCity   string `json:"receiver_city,omitempty"`

    // 收件人区县
    ReceiverArea   string `json:"receiver_area,omitempty"`

    // 收件人区县
    ReceiverAddress   string `json:"receiver_address,omitempty"`

    // 收件方邮编
    ReceiverZipCode   string `json:"receiver_zip_code,omitempty"`

    // 收件人姓名
    ReceiverName   string `json:"receiver_name,omitempty"`

    // 收件人昵称
    ReceiverNick   string `json:"receiver_nick,omitempty"`

    // 收件人手机，手机与电话必须有一值不为空
    ReceiverMobile   string `json:"receiver_mobile,omitempty"`

    // 收件人电话，手机与电话必须有一值不为空
    ReceiverPhone   string `json:"receiver_phone,omitempty"`

}
