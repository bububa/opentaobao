package btrip

// OpenInvoiceDeleteRq 
type OpenInvoiceDeleteRq struct {
    // 第三方企业id
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 第三方发票id
    ThirdPartId   string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
    // 商旅开放平台传2
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
}
