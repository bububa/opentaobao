package ottpay

// TvOrderResultDTO 
type TvOrderResultDTO struct {
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // orderNo
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    // cpOrderNo
    CpOrderNo   string `json:"cp_order_no,omitempty" xml:"cp_order_no,omitempty"`
    // qcodeUrl
    QcodeUrl   string `json:"qcode_url,omitempty" xml:"qcode_url,omitempty"`
    // 默认空
    VersionCode   string `json:"version_code,omitempty" xml:"version_code,omitempty"`
}
