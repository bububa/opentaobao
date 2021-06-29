package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流代货主查往来单位接口 APIRequest
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

func NewAlibabaAlihealthDrugKytListpartsByagentRequest() *AlibabaAlihealthDrugKytListpartsByagentRequest{
    return &AlibabaAlihealthDrugKytListpartsByagentRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.listparts.byagent"
}

func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetEntName() string {
    return r.entName
}

func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetRefPartnerId(refPartnerId string) error {
    r.refPartnerId = refPartnerId
    r.Set("ref_partner_id", refPartnerId)
    return nil
}

func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetRefPartnerId() string {
    return r.refPartnerId
}

func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetBeginDate() string {
    return r.beginDate
}

func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaAlihealthDrugKytListpartsByagentRequest) SetAgentRefEntId(agentRefEntId string) error {
    r.agentRefEntId = agentRefEntId
    r.Set("agent_ref_ent_id", agentRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytListpartsByagentRequest) GetAgentRefEntId() string {
    return r.agentRefEntId
}

