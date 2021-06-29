package eticket

// CertificateDTO 
type CertificateDTO struct {
    // attributes
    Attributes   *Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
    // availableNum
    AvailableNum   int64 `json:"available_num,omitempty" xml:"available_num,omitempty"`
    // bizType
    BizType   int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // codeStatus
    CodeStatus   int64 `json:"code_status,omitempty" xml:"code_status,omitempty"`
    // endTime
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // initialNum
    InitialNum   int64 `json:"initial_num,omitempty" xml:"initial_num,omitempty"`
    // lockedNum
    LockedNum   int64 `json:"locked_num,omitempty" xml:"locked_num,omitempty"`
    // outerId
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // qrCodeUrl
    QrCodeUrl   string `json:"qr_code_url,omitempty" xml:"qr_code_url,omitempty"`
    // startTime
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // usedNum
    UsedNum   int64 `json:"used_num,omitempty" xml:"used_num,omitempty"`
}
