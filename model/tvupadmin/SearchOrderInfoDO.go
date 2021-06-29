package tvupadmin

// SearchOrderInfoDO 
type SearchOrderInfoDO struct {

    // 创建时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 用户ID
    
    Uid   int64 `json:"uid,omitempty" xml:"uid,omitempty"`
    

    // 产品名称
    
    Subject   string `json:"subject,omitempty" xml:"subject,omitempty"`
    

    // 用户名
    
    UserLogonId   string `json:"user_logon_id,omitempty" xml:"user_logon_id,omitempty"`
    

    // 实际支付金额
    
    ActualPaidPrice   int64 `json:"actual_paid_price,omitempty" xml:"actual_paid_price,omitempty"`
    

    // 产品类型
    
    SubjectType   int64 `json:"subject_type,omitempty" xml:"subject_type,omitempty"`
    

    // 牌照方
    
    License   string `json:"license,omitempty" xml:"license,omitempty"`
    

    // 优酷用户ID
    
    YoukuUid   int64 `json:"youku_uid,omitempty" xml:"youku_uid,omitempty"`
    

    // 权益截止时间
    
    GmtVaildEnd   string `json:"gmt_vaild_end,omitempty" xml:"gmt_vaild_end,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 状态
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 优酷登录ID
    
    YoukuLogonId   string `json:"youku_logon_id,omitempty" xml:"youku_logon_id,omitempty"`
    

    // 权益开始时间
    
    GmtVaildStart   string `json:"gmt_vaild_start,omitempty" xml:"gmt_vaild_start,omitempty"`
    

    // 价格
    
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    

    // 订单号
    
    OrderNo   int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
    

    // 支付时间
    
    GmtPaySuccess   string `json:"gmt_pay_success,omitempty" xml:"gmt_pay_success,omitempty"`
    

    // 产品ID
    
    SubjectId   string `json:"subject_id,omitempty" xml:"subject_id,omitempty"`
    

    // 业务类型
    
    BizType   int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    

    // 客户端IP
    
    ClientIp   string `json:"client_ip,omitempty" xml:"client_ip,omitempty"`
    

    // 设备唯一标识
    
    Uuid   string `json:"uuid,omitempty" xml:"uuid,omitempty"`
    

}
