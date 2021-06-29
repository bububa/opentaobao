package drugtrace

// HaoxinqingDataDto 
type HaoxinqingDataDto struct {
    // 数量
    Number   int64 `json:"number,omitempty" xml:"number,omitempty"`
    // 业务类型， 1。咨询类  、2。购药类
    BussinessType   int64 `json:"bussiness_type,omitempty" xml:"bussiness_type,omitempty"`
    // 支付时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 药品名称
    DrugName   string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
    // 用户ID
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 用于唯一建，同一条信息uuid一样则更新
    Uuid   string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}
