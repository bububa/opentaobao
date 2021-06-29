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
    refEntId   string
    // 货主
    authRefUserId   string
    // 开始日期
    beginDate   string
    // 结束日期
    endDate   string
    // 发货企业
    partnerIdSend   string
    // 收货企业
    partnerIdRecv   string
    // 代理企业
    agentRefUserId   string
    // 当前页
    curPage   int64
    // 页大小
    pageSize   int64
    // 单据号码
    billCode   string
    // 单据类型  A : 所有  AI :入库    AO:出库
    billType   string
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
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetRefEntId() string {
    return r.refEntId
}
// AuthRefUserId Setter
// 货主
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetAuthRefUserId(authRefUserId string) error {
    r.authRefUserId = authRefUserId
    r.Set("auth_ref_user_id", authRefUserId)
    return nil
}

// AuthRefUserId Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetAuthRefUserId() string {
    return r.authRefUserId
}
// BeginDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetBeginDate() string {
    return r.beginDate
}
// EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetEndDate() string {
    return r.endDate
}
// PartnerIdSend Setter
// 发货企业
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetPartnerIdSend(partnerIdSend string) error {
    r.partnerIdSend = partnerIdSend
    r.Set("partner_id_send", partnerIdSend)
    return nil
}

// PartnerIdSend Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetPartnerIdSend() string {
    return r.partnerIdSend
}
// PartnerIdRecv Setter
// 收货企业
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetPartnerIdRecv(partnerIdRecv string) error {
    r.partnerIdRecv = partnerIdRecv
    r.Set("partner_id_recv", partnerIdRecv)
    return nil
}

// PartnerIdRecv Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetPartnerIdRecv() string {
    return r.partnerIdRecv
}
// AgentRefUserId Setter
// 代理企业
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}
// CurPage Setter
// 当前页
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetCurPage(curPage int64) error {
    r.curPage = curPage
    r.Set("cur_page", curPage)
    return nil
}

// CurPage Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetCurPage() int64 {
    return r.curPage
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetPageSize() int64 {
    return r.pageSize
}
// BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetBillCode() string {
    return r.billCode
}
// BillType Setter
// 单据类型  A : 所有  AI :入库    AO:出库
func (r *AlibabaAlihealthDrugKytSearchbillRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytSearchbillRequest) GetBillType() string {
    return r.billType
}
