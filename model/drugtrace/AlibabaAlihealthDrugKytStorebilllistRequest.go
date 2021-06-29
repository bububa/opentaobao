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
    _refEntId   string
    // 开始日期
    _startDate   string
    // 结束日期
    _endDate   string
    // 单据状态(A全部 1上传成功 3处理成功 4处理失败 )
    _billStatus   string
    // 页码
    _page   int64
    // 页数
    _pageSize   int64
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
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetRefEntId() string {
    return r._refEntId
}
// StartDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetEndDate() string {
    return r._endDate
}
// BillStatus Setter
// 单据状态(A全部 1上传成功 3处理成功 4处理失败 )
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetBillStatus(_billStatus string) error {
    r._billStatus = _billStatus
    r.Set("bill_status", _billStatus)
    return nil
}

// BillStatus Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetBillStatus() string {
    return r._billStatus
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 页数
func (r *AlibabaAlihealthDrugKytStorebilllistRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytStorebilllistRequest) GetPageSize() int64 {
    return r._pageSize
}
