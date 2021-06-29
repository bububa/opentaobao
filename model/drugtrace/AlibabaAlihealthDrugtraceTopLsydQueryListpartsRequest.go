package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
往来单位查询 API请求
alibaba.alihealth.drugtrace.top.lsyd.query.listparts

查询往来单位列表
*/
type AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest struct {
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

// 初始化AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest() *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest{
    return &AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.query.listparts"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) GetRefEntId() string {
    return r.refEntId
}
// EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) GetEntName() string {
    return r.entName
}
// RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetRefPartnerId(refPartnerId string) error {
    r.refPartnerId = refPartnerId
    r.Set("ref_partner_id", refPartnerId)
    return nil
}

// RefPartnerId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) GetRefPartnerId() string {
    return r.refPartnerId
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) GetPageSize() int64 {
    return r.pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) GetPage() int64 {
    return r.page
}
// BeginDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) GetBeginDate() string {
    return r.beginDate
}
// EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) GetEndDate() string {
    return r.endDate
}
