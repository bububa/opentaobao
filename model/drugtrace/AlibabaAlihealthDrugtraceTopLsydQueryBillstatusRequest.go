package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单据后处理状态查询 API请求
alibaba.alihealth.drugtrace.top.lsyd.query.billstatus

单据处理状态查询
*/
type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest struct {
    model.Params
    // 企业ID
    refEntId   string
    // 开始日期
    beginDate   string
    // 结束日期
    endDate   string
    // 单据类型 A：全部 AI：全部入库 AO：全部出库
    billType   string
    // 单据号
    billCode   string
    // 药品类型
    drugType   string
    // 状态  0, 上传成功     3, 处理成功     4, 处理失败
    dealStatus   string
    // 发货商
    fromUserId   string
    // 收货商
    toUserId   string
    // 代理商
    agentRefUserId   string
    // 页大小
    pageSize   int64
    // 页码
    page   int64
}

// 初始化AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest() *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest{
    return &AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.query.billstatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetRefEntId() string {
    return r.refEntId
}
// BeginDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetBeginDate() string {
    return r.beginDate
}
// EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetEndDate() string {
    return r.endDate
}
// BillType Setter
// 单据类型 A：全部 AI：全部入库 AO：全部出库
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetBillType() string {
    return r.billType
}
// BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetBillCode() string {
    return r.billCode
}
// DrugType Setter
// 药品类型
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetDrugType(drugType string) error {
    r.drugType = drugType
    r.Set("drug_type", drugType)
    return nil
}

// DrugType Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetDrugType() string {
    return r.drugType
}
// DealStatus Setter
// 状态  0, 上传成功     3, 处理成功     4, 处理失败
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetDealStatus(dealStatus string) error {
    r.dealStatus = dealStatus
    r.Set("deal_status", dealStatus)
    return nil
}

// DealStatus Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetDealStatus() string {
    return r.dealStatus
}
// FromUserId Setter
// 发货商
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetFromUserId() string {
    return r.fromUserId
}
// ToUserId Setter
// 收货商
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetToUserId(toUserId string) error {
    r.toUserId = toUserId
    r.Set("to_user_id", toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetToUserId() string {
    return r.toUserId
}
// AgentRefUserId Setter
// 代理商
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetPageSize() int64 {
    return r.pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetPage() int64 {
    return r.page
}
