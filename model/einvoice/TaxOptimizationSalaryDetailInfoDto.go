package einvoice

// TaxOptimizationSalaryDetailInfoDto 
type TaxOptimizationSalaryDetailInfoDto struct {
    // 明细金额
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    // 承包商编码
    ContractorCode   string `json:"contractor_code,omitempty" xml:"contractor_code,omitempty"`
    // 创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 明细id
    DetailId   string `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
    // 用户在业务平台的userid
    IdentificationInBelongingEmployer   string `json:"identification_in_belonging_employer,omitempty" xml:"identification_in_belonging_employer,omitempty"`
}
