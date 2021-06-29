package lstlogistics

// Content 
type Content struct {
    // 运单编号
    MailNo   string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
    // 物流编号
    LogisticsId   string `json:"logistics_id,omitempty" xml:"logistics_id,omitempty"`
    // 物流公司code
    CpCompanyCode   string `json:"cp_company_code,omitempty" xml:"cp_company_code,omitempty"`
    // 物流公司name
    CpCompanyName   string `json:"cp_company_name,omitempty" xml:"cp_company_name,omitempty"`
    // 子订单列表
    SubOrderIdList   []int64 `json:"sub_order_id_list,omitempty" xml:"sub_order_id_list>int64,omitempty"`
    // 描述
    StatusDesc   string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
    // 时间
    Time   string `json:"time,omitempty" xml:"time,omitempty"`
    // 描述
    StanderdDesc   string `json:"standerd_desc,omitempty" xml:"standerd_desc,omitempty"`
}
