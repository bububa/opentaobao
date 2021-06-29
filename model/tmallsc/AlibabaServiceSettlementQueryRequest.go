package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台结算单明细查询服务 API请求
alibaba.service.settlement.query

给服务商提供结算单明细查询功能
*/
type AlibabaServiceSettlementQueryRequest struct {
    model.Params
    // 账单查询开始时间。格式示例 2019-03-26 17:15:28
    gmtCreateStart   string
    // 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
    gmtCreateEnd   string
    // 当前页面，开始值为1
    currentPage   int64
    // 页面展示条数大小
    pageSize   int64
    // 工单ID
    workcardId   int64
    // 交易主订单号码
    parentTradeOrderId   int64
    // 服务单号
    serviceOrderId   int64
    // 交易实物订单号
    masterTradeOrderId   int64
    // 交易服务订单号
    serviceTradeOrderId   int64
    // 账单修改开始时间。
    gmtModifiedEnd   string
    // 账单修改结束时间，时间区间限制未15分钟。
    gmtModifiedStart   string
}

// 初始化AlibabaServiceSettlementQueryRequest对象
func NewAlibabaServiceSettlementQueryRequest() *AlibabaServiceSettlementQueryRequest{
    return &AlibabaServiceSettlementQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServiceSettlementQueryRequest) GetApiMethodName() string {
    return "alibaba.service.settlement.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServiceSettlementQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GmtCreateStart Setter
// 账单查询开始时间。格式示例 2019-03-26 17:15:28
func (r *AlibabaServiceSettlementQueryRequest) SetGmtCreateStart(gmtCreateStart string) error {
    r.gmtCreateStart = gmtCreateStart
    r.Set("gmt_create_start", gmtCreateStart)
    return nil
}

// GmtCreateStart Getter
func (r AlibabaServiceSettlementQueryRequest) GetGmtCreateStart() string {
    return r.gmtCreateStart
}
// GmtCreateEnd Setter
// 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
func (r *AlibabaServiceSettlementQueryRequest) SetGmtCreateEnd(gmtCreateEnd string) error {
    r.gmtCreateEnd = gmtCreateEnd
    r.Set("gmt_create_end", gmtCreateEnd)
    return nil
}

// GmtCreateEnd Getter
func (r AlibabaServiceSettlementQueryRequest) GetGmtCreateEnd() string {
    return r.gmtCreateEnd
}
// CurrentPage Setter
// 当前页面，开始值为1
func (r *AlibabaServiceSettlementQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaServiceSettlementQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// PageSize Setter
// 页面展示条数大小
func (r *AlibabaServiceSettlementQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaServiceSettlementQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// WorkcardId Setter
// 工单ID
func (r *AlibabaServiceSettlementQueryRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

// WorkcardId Getter
func (r AlibabaServiceSettlementQueryRequest) GetWorkcardId() int64 {
    return r.workcardId
}
// ParentTradeOrderId Setter
// 交易主订单号码
func (r *AlibabaServiceSettlementQueryRequest) SetParentTradeOrderId(parentTradeOrderId int64) error {
    r.parentTradeOrderId = parentTradeOrderId
    r.Set("parent_trade_order_id", parentTradeOrderId)
    return nil
}

// ParentTradeOrderId Getter
func (r AlibabaServiceSettlementQueryRequest) GetParentTradeOrderId() int64 {
    return r.parentTradeOrderId
}
// ServiceOrderId Setter
// 服务单号
func (r *AlibabaServiceSettlementQueryRequest) SetServiceOrderId(serviceOrderId int64) error {
    r.serviceOrderId = serviceOrderId
    r.Set("service_order_id", serviceOrderId)
    return nil
}

// ServiceOrderId Getter
func (r AlibabaServiceSettlementQueryRequest) GetServiceOrderId() int64 {
    return r.serviceOrderId
}
// MasterTradeOrderId Setter
// 交易实物订单号
func (r *AlibabaServiceSettlementQueryRequest) SetMasterTradeOrderId(masterTradeOrderId int64) error {
    r.masterTradeOrderId = masterTradeOrderId
    r.Set("master_trade_order_id", masterTradeOrderId)
    return nil
}

// MasterTradeOrderId Getter
func (r AlibabaServiceSettlementQueryRequest) GetMasterTradeOrderId() int64 {
    return r.masterTradeOrderId
}
// ServiceTradeOrderId Setter
// 交易服务订单号
func (r *AlibabaServiceSettlementQueryRequest) SetServiceTradeOrderId(serviceTradeOrderId int64) error {
    r.serviceTradeOrderId = serviceTradeOrderId
    r.Set("service_trade_order_id", serviceTradeOrderId)
    return nil
}

// ServiceTradeOrderId Getter
func (r AlibabaServiceSettlementQueryRequest) GetServiceTradeOrderId() int64 {
    return r.serviceTradeOrderId
}
// GmtModifiedEnd Setter
// 账单修改开始时间。
func (r *AlibabaServiceSettlementQueryRequest) SetGmtModifiedEnd(gmtModifiedEnd string) error {
    r.gmtModifiedEnd = gmtModifiedEnd
    r.Set("gmt_modified_end", gmtModifiedEnd)
    return nil
}

// GmtModifiedEnd Getter
func (r AlibabaServiceSettlementQueryRequest) GetGmtModifiedEnd() string {
    return r.gmtModifiedEnd
}
// GmtModifiedStart Setter
// 账单修改结束时间，时间区间限制未15分钟。
func (r *AlibabaServiceSettlementQueryRequest) SetGmtModifiedStart(gmtModifiedStart string) error {
    r.gmtModifiedStart = gmtModifiedStart
    r.Set("gmt_modified_start", gmtModifiedStart)
    return nil
}

// GmtModifiedStart Getter
func (r AlibabaServiceSettlementQueryRequest) GetGmtModifiedStart() string {
    return r.gmtModifiedStart
}
