package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询采购单信息 API请求
taobao.fenxiao.orders.get

分销商或供应商均可用此接口查询采购单信息（代销）； (发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。
*/
type TaobaoFenxiaoOrdersGetRequest struct {
    model.Params
    // 交易状态，不传默认查询所有采购单根据身份选择自身状态可选值: *供应商： WAIT_SELLER_SEND_GOODS(等待发货) WAIT_SELLER_CONFIRM_PAY(待确认收款) WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(已发货) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭) *分销商： WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(待收货确认) TRADE_FOR_PAY(已付款) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭)
    status   string
    // 起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
    startCreated   string
    // 结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
    endCreated   string
    // 可选值：trade_time_type(采购单按照成交时间范围查询),update_time_type(采购单按照更新时间范围查询)
    timeType   string
    // 页码。（大于0的整数。默认为1）
    pageNo   int64
    // 每页条数。（每页条数不超过50条）
    pageSize   int64
    // 采购单编号或分销流水号，若其它参数没传，则此参数必传。
    purchaseOrderId   int64
    // 多个字段用","分隔。<br/><br/>fields<br/>如果为空：返回所有采购单对象(purchase_orders)字段。<br/>如果不为空：返回指定采购单对象(purchase_orders)字段。<br/><br/>例1：<br/>sub_purchase_orders.tc_order_id 表示只返回tc_order_id <br/>例2：<br/>sub_purchase_orders表示只返回子采购单列表
    fields   string
    // 采购单下游买家订单id
    tcOrderId   int64
    // 渠道code，可批量 查询老供销：999
    channelCodes   []int64
    // 角色，供应商：2，分销商：1
    userRoleType   int64
    // 代销：1 经销：2 寄售（猫超自营寄售）：5 平台寄售：6
    tradeTypes   []int64
}

// 初始化TaobaoFenxiaoOrdersGetRequest对象
func NewTaobaoFenxiaoOrdersGetRequest() *TaobaoFenxiaoOrdersGetRequest{
    return &TaobaoFenxiaoOrdersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoOrdersGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.orders.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoOrdersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 交易状态，不传默认查询所有采购单根据身份选择自身状态可选值: *供应商： WAIT_SELLER_SEND_GOODS(等待发货) WAIT_SELLER_CONFIRM_PAY(待确认收款) WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(已发货) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭) *分销商： WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(待收货确认) TRADE_FOR_PAY(已付款) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭)
func (r *TaobaoFenxiaoOrdersGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetStatus() string {
    return r.status
}
// StartCreated Setter
// 起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
func (r *TaobaoFenxiaoOrdersGetRequest) SetStartCreated(startCreated string) error {
    r.startCreated = startCreated
    r.Set("start_created", startCreated)
    return nil
}

// StartCreated Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetStartCreated() string {
    return r.startCreated
}
// EndCreated Setter
// 结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
func (r *TaobaoFenxiaoOrdersGetRequest) SetEndCreated(endCreated string) error {
    r.endCreated = endCreated
    r.Set("end_created", endCreated)
    return nil
}

// EndCreated Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetEndCreated() string {
    return r.endCreated
}
// TimeType Setter
// 可选值：trade_time_type(采购单按照成交时间范围查询),update_time_type(采购单按照更新时间范围查询)
func (r *TaobaoFenxiaoOrdersGetRequest) SetTimeType(timeType string) error {
    r.timeType = timeType
    r.Set("time_type", timeType)
    return nil
}

// TimeType Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetTimeType() string {
    return r.timeType
}
// PageNo Setter
// 页码。（大于0的整数。默认为1）
func (r *TaobaoFenxiaoOrdersGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页条数。（每页条数不超过50条）
func (r *TaobaoFenxiaoOrdersGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// PurchaseOrderId Setter
// 采购单编号或分销流水号，若其它参数没传，则此参数必传。
func (r *TaobaoFenxiaoOrdersGetRequest) SetPurchaseOrderId(purchaseOrderId int64) error {
    r.purchaseOrderId = purchaseOrderId
    r.Set("purchase_order_id", purchaseOrderId)
    return nil
}

// PurchaseOrderId Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetPurchaseOrderId() int64 {
    return r.purchaseOrderId
}
// Fields Setter
// 多个字段用","分隔。<br/><br/>fields<br/>如果为空：返回所有采购单对象(purchase_orders)字段。<br/>如果不为空：返回指定采购单对象(purchase_orders)字段。<br/><br/>例1：<br/>sub_purchase_orders.tc_order_id 表示只返回tc_order_id <br/>例2：<br/>sub_purchase_orders表示只返回子采购单列表
func (r *TaobaoFenxiaoOrdersGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetFields() string {
    return r.fields
}
// TcOrderId Setter
// 采购单下游买家订单id
func (r *TaobaoFenxiaoOrdersGetRequest) SetTcOrderId(tcOrderId int64) error {
    r.tcOrderId = tcOrderId
    r.Set("tc_order_id", tcOrderId)
    return nil
}

// TcOrderId Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetTcOrderId() int64 {
    return r.tcOrderId
}
// ChannelCodes Setter
// 渠道code，可批量 查询老供销：999
func (r *TaobaoFenxiaoOrdersGetRequest) SetChannelCodes(channelCodes []int64) error {
    r.channelCodes = channelCodes
    r.Set("channel_codes", channelCodes)
    return nil
}

// ChannelCodes Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetChannelCodes() []int64 {
    return r.channelCodes
}
// UserRoleType Setter
// 角色，供应商：2，分销商：1
func (r *TaobaoFenxiaoOrdersGetRequest) SetUserRoleType(userRoleType int64) error {
    r.userRoleType = userRoleType
    r.Set("user_role_type", userRoleType)
    return nil
}

// UserRoleType Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetUserRoleType() int64 {
    return r.userRoleType
}
// TradeTypes Setter
// 代销：1 经销：2 寄售（猫超自营寄售）：5 平台寄售：6
func (r *TaobaoFenxiaoOrdersGetRequest) SetTradeTypes(tradeTypes []int64) error {
    r.tradeTypes = tradeTypes
    r.Set("trade_types", tradeTypes)
    return nil
}

// TradeTypes Getter
func (r TaobaoFenxiaoOrdersGetRequest) GetTradeTypes() []int64 {
    return r.tradeTypes
}
