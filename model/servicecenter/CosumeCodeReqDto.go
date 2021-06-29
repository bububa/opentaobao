package servicecenter

// CosumeCodeReqDTO 
type CosumeCodeReqDTO struct {
    // 业务id
    BizId   string `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
    // 业务类型,整车租赁传入：car_lease
    BizType   string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    // 车牌号
    CarNo   string `json:"car_no,omitempty" xml:"car_no,omitempty"`
    // 核销码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 身份证
    IdentityNo   string `json:"identity_no,omitempty" xml:"identity_no,omitempty"`
    // 门店id
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 门店名字
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // 车架号
    Vin   string `json:"vin,omitempty" xml:"vin,omitempty"`
}
