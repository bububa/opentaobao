package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单据处理状态查询 API请求
alibaba.alihealth.drug.kyt.searchstatus

单据处理状态查询
*/
type AlibabaAlihealthDrugKytSearchstatusRequest struct {
    model.Params
    // 企业ref_ent_id（货主企业的ref_ent_id）
    _refEntId   string
    // 开始日期（没有时分秒）
    _beginDate   string
    // 结束日期（没有时分秒）
    _endDate   string
    // 单据类型 A：全部 AI：全部入库 AO：全部出库
    _billType   string
    // 单据号（精确值，不支持模糊查询）
    _billCode   string
    // 药品类型
    _drugType   string
    // 状态  0, 处理中     3, 处理成功     4, 处理失败
    _dealStatus   string
    // 发货商
    _fromUserId   string
    // 收货商
    _toUserId   string
    // 代理商（第三方物流企业）
    _agentRefUserId   string
    // 页大小
    _pageSize   int64
    // 页码
    _page   int64
}

// 初始化AlibabaAlihealthDrugKytSearchstatusRequest对象
func NewAlibabaAlihealthDrugKytSearchstatusRequest() *AlibabaAlihealthDrugKytSearchstatusRequest{
    return &AlibabaAlihealthDrugKytSearchstatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.searchstatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ref_ent_id（货主企业的ref_ent_id）
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetRefEntId() string {
    return r._refEntId
}
// BeginDate Setter
// 开始日期（没有时分秒）
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetBeginDate(_beginDate string) error {
    r._beginDate = _beginDate
    r.Set("begin_date", _beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetBeginDate() string {
    return r._beginDate
}
// EndDate Setter
// 结束日期（没有时分秒）
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetEndDate() string {
    return r._endDate
}
// BillType Setter
// 单据类型 A：全部 AI：全部入库 AO：全部出库
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetBillType(_billType string) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetBillType() string {
    return r._billType
}
// BillCode Setter
// 单据号（精确值，不支持模糊查询）
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetBillCode() string {
    return r._billCode
}
// DrugType Setter
// 药品类型
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetDrugType(_drugType string) error {
    r._drugType = _drugType
    r.Set("drug_type", _drugType)
    return nil
}

// DrugType Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetDrugType() string {
    return r._drugType
}
// DealStatus Setter
// 状态  0, 处理中     3, 处理成功     4, 处理失败
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetDealStatus(_dealStatus string) error {
    r._dealStatus = _dealStatus
    r.Set("deal_status", _dealStatus)
    return nil
}

// DealStatus Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetDealStatus() string {
    return r._dealStatus
}
// FromUserId Setter
// 发货商
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetFromUserId() string {
    return r._fromUserId
}
// ToUserId Setter
// 收货商
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetToUserId(_toUserId string) error {
    r._toUserId = _toUserId
    r.Set("to_user_id", _toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetToUserId() string {
    return r._toUserId
}
// AgentRefUserId Setter
// 代理商（第三方物流企业）
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetAgentRefUserId(_agentRefUserId string) error {
    r._agentRefUserId = _agentRefUserId
    r.Set("agent_ref_user_id", _agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetAgentRefUserId() string {
    return r._agentRefUserId
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetPageSize() int64 {
    return r._pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytSearchstatusRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytSearchstatusRequest) GetPage() int64 {
    return r._page
}
