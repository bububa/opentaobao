package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
往来单位查询 APIRequest
alibaba.alihealth.drugtrace.top.yljg.query.listparts

查询往来单位列表
*/
type AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest struct {
    model.Params

    // 企业唯一标识
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

}

func NewAlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest() *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest{
    return &AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.query.listparts"
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) GetEntName() string {
    return r.entName
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetRefPartnerId(refPartnerId string) error {
    r.refPartnerId = refPartnerId
    r.Set("ref_partner_id", refPartnerId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) GetRefPartnerId() string {
    return r.refPartnerId
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) GetBeginDate() string {
    return r.beginDate
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) GetEndDate() string {
    return r.endDate
}

