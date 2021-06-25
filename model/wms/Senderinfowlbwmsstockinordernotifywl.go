package wms

// Senderinfowlbwmsstockinordernotifywl 
type Senderinfowlbwmsstockinordernotifywl struct {

    // 发件方邮编
    SenderZipCode   string `json:"sender_zip_code,omitempty"`

    // 发件方镇
    SenderTown   string `json:"sender_town,omitempty"`

    // 发件方地址
    SenderAddress   string `json:"sender_address,omitempty"`

    // 发件方名称，销退场景下填写买家名称； 调拨场景下填写调拨出仓库名称；
    SenderName   string `json:"sender_name,omitempty"`

    // 发件方省份
    SenderProvince   string `json:"sender_province,omitempty"`

    // 发件方区县
    SenderArea   string `json:"sender_area,omitempty"`

    // 发件方编码，销退场景下填写买家nick，旺旺号等； 调拨场景下填写调拨出仓库编码；
    SenderCode   string `json:"sender_code,omitempty"`

    // 发件方城市
    SenderCity   string `json:"sender_city,omitempty"`

    // 发件方手机
    SenderMobile   string `json:"sender_mobile,omitempty"`

    // 发件方电话
    SenderPhone   string `json:"sender_phone,omitempty"`

}
