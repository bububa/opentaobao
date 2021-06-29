package traveltrade

// VoucherInfoDTO 
type VoucherInfoDTO struct {
    // 用户短信会收到的确认号
    ConfirmCode   string `json:"confirm_code,omitempty" xml:"confirm_code,omitempty"`
    // 凭证使用次数
    UsedQuantity   int64 `json:"used_quantity,omitempty" xml:"used_quantity,omitempty"`
    // 凭证使用时间，格式:yyyy-MM-dd HH:mm:ss
    UsedDate   string `json:"used_date,omitempty" xml:"used_date,omitempty"`
}
