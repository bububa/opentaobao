package txcs

// VerificationBillResponseDto 
type VerificationBillResponseDto struct {
    // 核销单号
    VerificationNo   string `json:"verification_no,omitempty" xml:"verification_no,omitempty"`
    // 核销日期
    VerifyDate   string `json:"verify_date,omitempty" xml:"verify_date,omitempty"`
}
