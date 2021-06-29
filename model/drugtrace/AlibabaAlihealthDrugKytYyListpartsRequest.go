package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询往来单位 APIRequest
alibaba.alihealth.drug.kyt.yy.listparts

查询往来单位列表
*/
type AlibabaAlihealthDrugKytYyListpartsRequest struct {
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

func NewAlibabaAlihealthDrugKytYyListpartsRequest() *AlibabaAlihealthDrugKytYyListpartsRequest{
    return &AlibabaAlihealthDrugKytYyListpartsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytYyListpartsRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.listparts"
}

func (r AlibabaAlihealthDrugKytYyListpartsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytYyListpartsRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyListpartsRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytYyListpartsRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugKytYyListpartsRequest) GetEntName() string {
    return r.entName
}

func (r *AlibabaAlihealthDrugKytYyListpartsRequest) SetRefPartnerId(refPartnerId string) error {
    r.refPartnerId = refPartnerId
    r.Set("ref_partner_id", refPartnerId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyListpartsRequest) GetRefPartnerId() string {
    return r.refPartnerId
}

func (r *AlibabaAlihealthDrugKytYyListpartsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugKytYyListpartsRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthDrugKytYyListpartsRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugKytYyListpartsRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthDrugKytYyListpartsRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

func (r AlibabaAlihealthDrugKytYyListpartsRequest) GetBeginDate() string {
    return r.beginDate
}

func (r *AlibabaAlihealthDrugKytYyListpartsRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthDrugKytYyListpartsRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaAlihealthDrugKytYyListpartsRequest) SetAgentRefEntId(agentRefEntId string) error {
    r.agentRefEntId = agentRefEntId
    r.Set("agent_ref_ent_id", agentRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyListpartsRequest) GetAgentRefEntId() string {
    return r.agentRefEntId
}

