package wdk

// ErpFeedbackDto 
type ErpFeedbackDto struct {

    // 单据号
    BizOrderCode   string `json:"biz_order_code,omitempty"`

    // 反馈日期，反馈结果产生的时间
    FeedbackDate   string `json:"feedback_date,omitempty"`

    // 商品明细列表（子表）商品明细列表（子表）
    ItemList   []ErpFeedbackDetailDto `json:"item_list,omitempty"`

    // 原始单据号，需要提供原始配货申请单据号
    OriginalBillCode   string `json:"original_bill_code,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 店仓code，指的是反馈对象，对应一个物理店或仓编码
    WarehouseCode   string `json:"warehouse_code,omitempty"`

}
