package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
药店查询往来单位 APIRequest
alibaba.alihealth.drug.kyt.smyx.listparts

查询往来单位列表
*/
type AlibabaAlihealthDrugKytSmyxListpartsRequest struct {
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

func NewAlibabaAlihealthDrugKytSmyxListpartsRequest() *AlibabaAlihealthDrugKytSmyxListpartsRequest{
    return &AlibabaAlihealthDrugKytSmyxListpartsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytSmyxListpartsRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.smyx.listparts"
}

func (r AlibabaAlihealthDrugKytSmyxListpartsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytSmyxListpartsRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytSmyxListpartsRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytSmyxListpartsRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugKytSmyxListpartsRequest) GetEntName() string {
    return r.entName
}

func (r *AlibabaAlihealthDrugKytSmyxListpartsRequest) SetRefPartnerId(refPartnerId string) error {
    r.refPartnerId = refPartnerId
    r.Set("ref_partner_id", refPartnerId)
    return nil
}

func (r AlibabaAlihealthDrugKytSmyxListpartsRequest) GetRefPartnerId() string {
    return r.refPartnerId
}

func (r *AlibabaAlihealthDrugKytSmyxListpartsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugKytSmyxListpartsRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthDrugKytSmyxListpartsRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugKytSmyxListpartsRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthDrugKytSmyxListpartsRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

func (r AlibabaAlihealthDrugKytSmyxListpartsRequest) GetBeginDate() string {
    return r.beginDate
}

func (r *AlibabaAlihealthDrugKytSmyxListpartsRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthDrugKytSmyxListpartsRequest) GetEndDate() string {
    return r.endDate
}

