package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售端平台单据查询 APIRequest
alibaba.alihealth.drug.kyt.storebilllist

零售端平台单据查询
*/
type AlibabaAlihealthDrugKytStorebilllistRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

    // 开始日期
    startDate   string 

    // 结束日期
    endDate   string 

    // 单据状态(A全部 1上传成功 3处理成功 4处理失败 )
    billStatus   string 

    // 页码
    page   int64 

    // 页数
    pageSize   int64 

}

func NewAlibabaAlihealthDrugKytStorebilllistRequest() *AlibabaAlihealthDrugKytStorebilllistRequest{
    return &AlibabaAlihealthDrugKytStorebilllistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.storebilllist"
}

func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetStartDate() string {
    return r.startDate
}

func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetBillStatus(billStatus string) error {
    r.billStatus = billStatus
    r.Set("bill_status", billStatus)
    return nil
}

func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetBillStatus() string {
    return r.billStatus
}

func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetPageSize() int64 {
    return r.pageSize
}

