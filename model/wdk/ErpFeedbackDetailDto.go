package wdk

// ErpFeedbackDetailDto 
type ErpFeedbackDetailDto struct {

    // 数量
    Count   string `json:"count,omitempty"`

    // 商品code，盒马系统中的商品编码
    ItemCode   string `json:"item_code,omitempty"`

    // 处理意见，包括三种处理情况：允许减库存，允许线下换货，驳回质量反馈要求(1,减库存 2,换货 0,驳回 3,加库存)
    Operation   int64 `json:"operation,omitempty"`

    // 备注
    Reason   string `json:"reason,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 部门
    DeptCode   string `json:"dept_code,omitempty"`

    // 指定库位，可空
    CabinetCode   string `json:"cabinet_code,omitempty"`

}
