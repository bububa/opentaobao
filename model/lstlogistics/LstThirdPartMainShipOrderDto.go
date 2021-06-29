package lstlogistics

// LstThirdPartMainShipOrderDto 
type LstThirdPartMainShipOrderDto struct {

    // 货品列表
    
    Details   []LstThirdPartDetailShipOrderDto `json:"details,omitempty" xml:"details,omitempty"`
    

    // 签收时间
    
    SignTime   string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
    

    // 订单
    
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // 发货单状态：NEW --->新建，LOAD_WAIT--->待装车，LOAD_SUCCESS--->已装车，SIGN_SUCCESS--->签收，SIGN_PART_SUCCESS--->部分签收，SIGN_FAILED--->拒签，CANCEL--->取消
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 外部订单ID
    
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    

    // 发货单更新时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 发货单生成时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 发货单ID
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 装车时间
    
    LoadTime   string `json:"load_time,omitempty" xml:"load_time,omitempty"`
    

}
