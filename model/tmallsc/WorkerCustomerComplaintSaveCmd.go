package tmallsc

// WorkerCustomerComplaintSaveCmd 
type WorkerCustomerComplaintSaveCmd struct {
    // 商家昵称
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    // 工单号
    WorkcardId   int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
    // 外部单号
    OutId   string `json:"out_id,omitempty" xml:"out_id,omitempty"`
    // 外部单号类型，1：服务商系统工单号，2：CCO系统工单号
    OutType   int64 `json:"out_type,omitempty" xml:"out_type,omitempty"`
    // 工人身份证号码
    IdNumber   string `json:"id_number,omitempty" xml:"id_number,omitempty"`
    // 客诉类型:1:服务态度,2:服务质量,3:服务时效,4:服务乱收费,5:服务不规范
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 客诉来源：1：服务商  2：消费者 3：商家
    Source   int64 `json:"source,omitempty" xml:"source,omitempty"`
    // 处理措施，解决方案文字描述
    TreatmentMeasures   string `json:"treatment_measures,omitempty" xml:"treatment_measures,omitempty"`
    // 幂等键，请求唯一标识
    IdempotentId   string `json:"idempotent_id,omitempty" xml:"idempotent_id,omitempty"`
    // 客诉发起时间
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 整改完成时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 是否成立：1：成立，0：不成立
    Established   int64 `json:"established,omitempty" xml:"established,omitempty"`
    // 扩展字段，json对象字符串，如：{"serviceCode":"浴霸安装"}
    ExtendInfo   string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
}
