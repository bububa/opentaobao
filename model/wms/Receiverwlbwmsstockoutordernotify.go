package wms

// Receiverwlbwmsstockoutordernotify 
type Receiverwlbwmsstockoutordernotify struct {
    // 收件方城市
    ReceiverCity   string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
    // 收件方区县
    ReceiverArea   string `json:"receiver_area,omitempty" xml:"receiver_area,omitempty"`
    // 收件方手机
    ReceiverMobile   string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
    // 收件方省份
    ReceiverProvince   string `json:"receiver_province,omitempty" xml:"receiver_province,omitempty"`
    // 收件方名称
    ReceiverName   string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
    // 退供场景ECP填充供应商编码，调拨出库单ECP填充调拨入仓库编码, B2B出库单填写分销商ID(无分销ID的null)
    ReceiverCode   string `json:"receiver_code,omitempty" xml:"receiver_code,omitempty"`
    // 收件方邮编
    ReceiverZipCode   string `json:"receiver_zip_code,omitempty" xml:"receiver_zip_code,omitempty"`
    // 收件方地址
    ReceiverAddress   string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
    // 收件方电话
    ReceiverPhone   string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
}
