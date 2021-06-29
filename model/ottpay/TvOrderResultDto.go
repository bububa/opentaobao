package ottpay

// TvOrderResultDto 
type TvOrderResultDto struct {
    // qcodeUrl
    QcodeUrl   string `json:"qcode_url,omitempty" xml:"qcode_url,omitempty"`
    // 版本号
    VersionCode   string `json:"version_code,omitempty" xml:"version_code,omitempty"`
    // cp本次订单号
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    // cp原始订单号
    CpOrderNo   string `json:"cp_order_no,omitempty" xml:"cp_order_no,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
