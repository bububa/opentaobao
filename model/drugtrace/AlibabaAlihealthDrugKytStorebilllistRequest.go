package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售端平台单据查询 API请求
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

// 初始化AlibabaAlihealthDrugKytStorebilllistRequest对象
func NewAlibabaAlihealthDrugKytStorebilllistRequest() *AlibabaAlihealthDrugKytStorebilllistRequest{
    return &AlibabaAlihealthDrugKytStorebilllistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.storebilllist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetRefEntId() string {
    return r.refEntId
}
// StartDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetEndDate() string {
    return r.endDate
}
// BillStatus Setter
// 单据状态(A全部 1上传成功 3处理成功 4处理失败 )
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetBillStatus(billStatus string) error {
    r.billStatus = billStatus
    r.Set("bill_status", billStatus)
    return nil
}

// BillStatus Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetBillStatus() string {
    return r.billStatus
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetPage() int64 {
    return r.page
}
// PageSize Setter
// 页数
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetPageSize() int64 {
    return r.pageSize
}
