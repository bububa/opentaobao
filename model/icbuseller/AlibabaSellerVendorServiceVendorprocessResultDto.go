package icbuseller

// AlibabaSellerVendorServiceVendorprocessResultDto 
type AlibabaSellerVendorServiceVendorprocessResultDto struct {
    // 异常说明
    ExecDescription   string `json:"exec_description,omitempty" xml:"exec_description,omitempty"`
    // 状态码
    ReturnCode   int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
    // 返回集合
    List   []VendorMerchantRecordBaseDto `json:"list,omitempty" xml:"list>vendor_merchant_record_base_dto,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
