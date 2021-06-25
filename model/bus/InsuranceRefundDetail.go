package bus

// InsuranceRefundDetail 
type InsuranceRefundDetail struct {

    // 保险退款总金额
    InsurePrice   int64 `json:"insure_price,omitempty"`

    // 乘客证件号
    RiderCertNumber   string `json:"rider_cert_number,omitempty"`

    // 证件类型：01：身份证，02：护照，03：港澳通行证，04：台湾通行证，05：港澳往返内地通行证，06：台湾往返内地通行证，07：港澳居民居住证
    RiderCertType   string `json:"rider_cert_type,omitempty"`

    // 乘客姓名
    RiderName   string `json:"rider_name,omitempty"`

    // 退保信息
    TvmInsuranceInfos   []TvmInsuranceInfo `json:"tvm_insurance_infos,omitempty"`

}
