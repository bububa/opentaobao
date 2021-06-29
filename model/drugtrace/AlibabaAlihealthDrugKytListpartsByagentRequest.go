package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流代货主查往来单位接口 API请求
alibaba.alihealth.drug.kyt.listparts.byagent

代理企业查询往来单位列表
*/
type AlibabaAlihealthDrugKytListpartsByagentRequest struct {
    model.Params
    // 企业唯一标识（货主企业）
    refEntId   string
    // 企业名称
    entName   string
    // 企业自定义编号
    refPartnerId   string
    // 页大小
    pageSize   int64
    // 页码
    page   int64
    // 开始时间
    beginDate   string
    // 结束时间
    endDate   string
    // 代理企业唯一标识（物流企业）
    agentRefEntId   string
}

// 初始化AlibabaAlihealthDrugKytListpartsByagentRequest对象
func NewAlibabaAlihealthDrugKytListpartsByagentRequest() *AlibabaAlihealthDrugKytListpartsByagentRequest{
    return &AlibabaAlihealthDrugKytListpartsByagentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.listparts.byagent"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业唯一标识（货主企业）
func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetRefEntId() string {
    return r.refEntId
}
// EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetEntName() string {
    return r.entName
}
// RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetRefPartnerId(refPartnerId string) error {
    r.refPartnerId = refPartnerId
    r.Set("ref_partner_id", refPartnerId)
    return nil
}

// RefPartnerId Getter
func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetRefPartnerId() string {
    return r.refPartnerId
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetPageSize() int64 {
    return r.pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetPage() int64 {
    return r.page
}
// BeginDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetBeginDate() string {
    return r.beginDate
}
// EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetEndDate() string {
    return r.endDate
}
// AgentRefEntId Setter
// 代理企业唯一标识（物流企业）
func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetAgentRefEntId(agentRefEntId string) error {
    r.agentRefEntId = agentRefEntId
    r.Set("agent_ref_ent_id", agentRefEntId)
    return nil
}

// AgentRefEntId Getter
func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetAgentRefEntId() string {
    return r.agentRefEntId
}
