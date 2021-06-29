package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多融查询一个企业的往来单位 APIRequest
alibaba.alihealth.drug.kyt.dr.listparts

查询往来单位列表
*/
type AlibabaAlihealthDrugKytDrListpartsRequest struct {
    model.Params

    // 企业ID
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

func NewAlibabaAlihealthDrugKytDrListpartsRequest() *AlibabaAlihealthDrugKytDrListpartsRequest{
    return &AlibabaAlihealthDrugKytDrListpartsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrListpartsRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.listparts"
}

func (r AlibabaAlihealthDrugKytDrListpartsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrListpartsRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrListpartsRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDrListpartsRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugKytDrListpartsRequest) GetEntName() string {
    return r.entName
}

func (r *AlibabaAlihealthDrugKytDrListpartsRequest) SetRefPartnerId(refPartnerId string) error {
    r.refPartnerId = refPartnerId
    r.Set("ref_partner_id", refPartnerId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrListpartsRequest) GetRefPartnerId() string {
    return r.refPartnerId
}

func (r *AlibabaAlihealthDrugKytDrListpartsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugKytDrListpartsRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthDrugKytDrListpartsRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugKytDrListpartsRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthDrugKytDrListpartsRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

func (r AlibabaAlihealthDrugKytDrListpartsRequest) GetBeginDate() string {
    return r.beginDate
}

func (r *AlibabaAlihealthDrugKytDrListpartsRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthDrugKytDrListpartsRequest) GetEndDate() string {
    return r.endDate
}

