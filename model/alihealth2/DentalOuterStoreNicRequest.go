package alihealth2

// DentalOuterStoreNicRequest 
type DentalOuterStoreNicRequest struct {
    // 失效时间
    ExpireTime   string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
    // 签约时间
    SignTime   string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
    // 门店ID
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 签约状态(0 未签约，1 已签约)
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}
