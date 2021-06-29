package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过时间段批量查询入出库单信息 API请求
alibaba.alihealth.drug.kyt.searchbill

通过时间段批量查询入出库单信息
*/
type AlibabaAlihealthDrugKytSearchbillRequest struct {
    model.Params
    // 企业标识
    _refEntId   string
    // 货主
    _authRefUserId   string
    // 开始日期
    _beginDate   string
    // 结束日期
    _endDate   string
    // 发货企业
    _partnerIdSend   string
    // 收货企业
    _partnerIdRecv   string
    // 代理企业
    _agentRefUserId   string
    // 当前页
    _curPage   int64
    // 页大小
    _pageSize   int64
    // 单据号码
    _billCode   string
    // 单据类型  A : 所有  AI :入库    AO:出库
    _billType   string
}

// 初始化AlibabaAlihealthDrugKytSearchbillRequest对象
func NewAlibabaAlihealthDrugKytSearchbillRequest() *AlibabaAlihealthDrugKytSearchbillRequest{
    return &AlibabaAlihealthDrugKytSearchbillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.searchbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业标识
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetRefEntId() string {
    return r._refEntId
}
// AuthRefUserId Setter
// 货主
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetAuthRefUserId(_authRefUserId string) error {
    r._authRefUserId = _authRefUserId
    r.Set("auth_ref_user_id", _authRefUserId)
    return nil
}

// AuthRefUserId Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetAuthRefUserId() string {
    return r._authRefUserId
}
// BeginDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetBeginDate(_beginDate string) error {
    r._beginDate = _beginDate
    r.Set("begin_date", _beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetBeginDate() string {
    return r._beginDate
}
// EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetEndDate() string {
    return r._endDate
}
// PartnerIdSend Setter
// 发货企业
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetPartnerIdSend(_partnerIdSend string) error {
    r._partnerIdSend = _partnerIdSend
    r.Set("partner_id_send", _partnerIdSend)
    return nil
}

// PartnerIdSend Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetPartnerIdSend() string {
    return r._partnerIdSend
}
// PartnerIdRecv Setter
// 收货企业
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetPartnerIdRecv(_partnerIdRecv string) error {
    r._partnerIdRecv = _partnerIdRecv
    r.Set("partner_id_recv", _partnerIdRecv)
    return nil
}

// PartnerIdRecv Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetPartnerIdRecv() string {
    return r._partnerIdRecv
}
// AgentRefUserId Setter
// 代理企业
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetAgentRefUserId(_agentRefUserId string) error {
    r._agentRefUserId = _agentRefUserId
    r.Set("agent_ref_user_id", _agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetAgentRefUserId() string {
    return r._agentRefUserId
}
// CurPage Setter
// 当前页
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetCurPage(_curPage int64) error {
    r._curPage = _curPage
    r.Set("cur_page", _curPage)
    return nil
}

// CurPage Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetCurPage() int64 {
    return r._curPage
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetPageSize() int64 {
    return r._pageSize
}
// BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetBillCode() string {
    return r._billCode
}
// BillType Setter
// 单据类型  A : 所有  AI :入库    AO:出库
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetBillType(_billType string) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetBillType() string {
    return r._billType
}
