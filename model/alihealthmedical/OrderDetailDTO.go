package alihealthmedical

// OrderDetailDTO 
type OrderDetailDTO struct {
    // 商品类型
    ItemType   string `json:"item_type,omitempty" xml:"item_type,omitempty"`
    // 结算日期，格式必须为：20200711
    CloseAccountTime   string `json:"close_account_time,omitempty" xml:"close_account_time,omitempty"`
    // 短评标签名字, 使用英文逗号分隔
    LabelNames   string `json:"label_names,omitempty" xml:"label_names,omitempty"`
    // 订单接诊时间
    DiagnosingTime   string `json:"diagnosing_time,omitempty" xml:"diagnosing_time,omitempty"`
    // 订单id
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 就诊人UUID，最长64个字符
    PatientId   string `json:"patient_id,omitempty" xml:"patient_id,omitempty"`
    // 订单下单时间
    OrderCreateTime   string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
    // 订单状态：WAIT_DIAGNOSE，  DIAGNOSING，  REFUNDED，  DIAGNOSED，  THIRD_SELLER_PAID
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 就诊人信息串
    MedicalInfo   *MedicalInfoDTO `json:"medical_info,omitempty" xml:"medical_info,omitempty"`
    // 评价综合得分：取值 1 - 5
    Score   int64 `json:"score,omitempty" xml:"score,omitempty"`
    // 商品id
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 医生名字
    DoctorName   string `json:"doctor_name,omitempty" xml:"doctor_name,omitempty"`
    // 评价时间
    ReviewCreateTime   string `json:"review_create_time,omitempty" xml:"review_create_time,omitempty"`
    // 商品总价，单位为元
    TotalFee   string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
    // 医生UUID
    DoctorId   string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
    // 订单终止时间
    OrderFinishTime   string `json:"order_finish_time,omitempty" xml:"order_finish_time,omitempty"`
    // 评价内容
    Comment   string `json:"comment,omitempty" xml:"comment,omitempty"`
    // 会话id
    SessionId   string `json:"session_id,omitempty" xml:"session_id,omitempty"`
}
