package eticket

// CertificateDto 
type CertificateDto struct {

    // attributes
    Attributes   *Attributes `json:"attributes,omitempty"`

    // availableNum
    AvailableNum   int64 `json:"available_num,omitempty"`

    // bizType
    BizType   int64 `json:"biz_type,omitempty"`

    // code
    Code   string `json:"code,omitempty"`

    // codeStatus
    CodeStatus   int64 `json:"code_status,omitempty"`

    // endTime
    EndTime   string `json:"end_time,omitempty"`

    // initialNum
    InitialNum   int64 `json:"initial_num,omitempty"`

    // lockedNum
    LockedNum   int64 `json:"locked_num,omitempty"`

    // outerId
    OuterId   string `json:"outer_id,omitempty"`

    // qrCodeUrl
    QrCodeUrl   string `json:"qr_code_url,omitempty"`

    // startTime
    StartTime   string `json:"start_time,omitempty"`

    // usedNum
    UsedNum   int64 `json:"used_num,omitempty"`

}
