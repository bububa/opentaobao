package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗企业上传单据后处理状态查询 API请求
alibaba.alihealth.drug.kyt.dr.searchstatus

单据处理状态查询
*/
type AlibabaAlihealthDrugKytDrSearchstatusRequest struct {
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

// 初始化AlibabaAlihealthDrugKytDrSearchstatusRequest对象
func NewAlibabaAlihealthDrugKytDrSearchstatusRequest() *AlibabaAlihealthDrugKytDrSearchstatusRequest{
    return &AlibabaAlihealthDrugKytDrSearchstatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.searchstatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetRefEntId() string {
    return r.refEntId
}
// BeginDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetBeginDate() string {
    return r.beginDate
}
// EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetEndDate() string {
    return r.endDate
}
// BillType Setter
// 单据类型 A：全部 AI：全部入库 AO：全部出库
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetBillType() string {
    return r.billType
}
// BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetBillCode() string {
    return r.billCode
}
// DrugType Setter
// 药品类型
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetDrugType(drugType string) error {
    r.drugType = drugType
    r.Set("drug_type", drugType)
    return nil
}

// DrugType Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetDrugType() string {
    return r.drugType
}
// DealStatus Setter
// 状态  0, 上传成功     3, 处理成功     4, 处理失败
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetDealStatus(dealStatus string) error {
    r.dealStatus = dealStatus
    r.Set("deal_status", dealStatus)
    return nil
}

// DealStatus Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetDealStatus() string {
    return r.dealStatus
}
// FromUserId Setter
// 发货商
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetFromUserId() string {
    return r.fromUserId
}
// ToUserId Setter
// 收货商
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetToUserId(toUserId string) error {
    r.toUserId = toUserId
    r.Set("to_user_id", toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetToUserId() string {
    return r.toUserId
}
// AgentRefUserId Setter
// 代理商
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetPageSize() int64 {
    return r.pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytDrSearchstatusRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytDrSearchstatusRequest) GetPage() int64 {
    return r.page
}
