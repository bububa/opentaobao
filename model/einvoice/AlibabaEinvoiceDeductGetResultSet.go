package einvoice

// AlibabaEinvoiceDeductGetResultSet 
type AlibabaEinvoiceDeductGetResultSet struct {
    // errorMessage
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // result
    Result   *AlibabaEinvoiceDeductGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
    // totalCount
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 业务日期
    BizDate   string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
    // 实际扣减
    Deduct   int64 `json:"deduct,omitempty" xml:"deduct,omitempty"`
    // 应扣减
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    // 税号
    PackageRegisterNo   string `json:"package_register_no,omitempty" xml:"package_register_no,omitempty"`
}
