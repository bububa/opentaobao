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
    _refEntId   string
    // 开始日期
    _beginDate   string
    // 结束日期
    _endDate   string
    // 单据类型 A：全部 AI：全部入库 AO：全部出库
    _billType   string
    // 单据号
    _billCode   string
    // 药品类型
    _drugType   string
    // 状态  0, 上传成功     3, 处理成功     4, 处理失败
    _dealStatus   string
    // 发货商
    _fromUserId   string
    // 收货商
    _toUserId   string
    // 代理商
    _agentRefUserId   string
    // 页大小
    _pageSize   int64
    // 页码
    _page   int64
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
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetRefEntId() string {
    return r._refEntId
}
// BeginDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetBeginDate(_beginDate string) error {
    r._beginDate = _beginDate
    r.Set("begin_date", _beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetBeginDate() string {
    return r._beginDate
}
// EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetEndDate() string {
    return r._endDate
}
// BillType Setter
// 单据类型 A：全部 AI：全部入库 AO：全部出库
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetBillType(_billType string) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetBillType() string {
    return r._billType
}
// BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetBillCode() string {
    return r._billCode
}
// DrugType Setter
// 药品类型
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetDrugType(_drugType string) error {
    r._drugType = _drugType
    r.Set("drug_type", _drugType)
    return nil
}

// DrugType Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetDrugType() string {
    return r._drugType
}
// DealStatus Setter
// 状态  0, 上传成功     3, 处理成功     4, 处理失败
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetDealStatus(_dealStatus string) error {
    r._dealStatus = _dealStatus
    r.Set("deal_status", _dealStatus)
    return nil
}

// DealStatus Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetDealStatus() string {
    return r._dealStatus
}
// FromUserId Setter
// 发货商
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetFromUserId() string {
    return r._fromUserId
}
// ToUserId Setter
// 收货商
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetToUserId(_toUserId string) error {
    r._toUserId = _toUserId
    r.Set("to_user_id", _toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetToUserId() string {
    return r._toUserId
}
// AgentRefUserId Setter
// 代理商
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetAgentRefUserId(_agentRefUserId string) error {
    r._agentRefUserId = _agentRefUserId
    r.Set("agent_ref_user_id", _agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetAgentRefUserId() string {
    return r._agentRefUserId
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetPageSize() int64 {
    return r._pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) GetPage() int64 {
    return r._page
}
