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
type TaobaoFenxiaoOrdersGetAPIRequest struct {
    model.Params
    // 交易状态，不传默认查询所有采购单根据身份选择自身状态可选值: *供应商： WAIT_SELLER_SEND_GOODS(等待发货) WAIT_SELLER_CONFIRM_PAY(待确认收款) WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(已发货) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭) *分销商： WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(待收货确认) TRADE_FOR_PAY(已付款) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭)
    _status   string
    // 起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
    _startCreated   string
    // 结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
    _endCreated   string
    // 可选值：trade_time_type(采购单按照成交时间范围查询),update_time_type(采购单按照更新时间范围查询)
    _timeType   string
    // 页码。（大于0的整数。默认为1）
    _pageNo   int64
    // 每页条数。（每页条数不超过50条）
    _pageSize   int64
    // 采购单编号或分销流水号，若其它参数没传，则此参数必传。
    _purchaseOrderId   int64
    // 多个字段用","分隔。<br/><br/>fields<br/>如果为空：返回所有采购单对象(purchase_orders)字段。<br/>如果不为空：返回指定采购单对象(purchase_orders)字段。<br/><br/>例1：<br/>sub_purchase_orders.tc_order_id 表示只返回tc_order_id <br/>例2：<br/>sub_purchase_orders表示只返回子采购单列表
    _fields   string
    // 采购单下游买家订单id
    _tcOrderId   int64
    // 渠道code，可批量 查询老供销：999
    _channelCodes   []int64
    // 角色，供应商：2，分销商：1
    _userRoleType   int64
    // 代销：1 经销：2 寄售（猫超自营寄售）：5 平台寄售：6
    _tradeTypes   []int64
}

// 初始化TaobaoFenxiaoOrdersGetAPIRequest对象
func NewTaobaoFenxiaoOrdersGetRequest() *TaobaoFenxiaoOrdersGetAPIRequest{
    return &TaobaoFenxiaoOrdersGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.orders.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 交易状态，不传默认查询所有采购单根据身份选择自身状态可选值: *供应商： WAIT_SELLER_SEND_GOODS(等待发货) WAIT_SELLER_CONFIRM_PAY(待确认收款) WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(已发货) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭) *分销商： WAIT_BUYER_PAY(等待付款) WAIT_BUYER_CONFIRM_GOODS(待收货确认) TRADE_FOR_PAY(已付款) TRADE_REFUNDING(退款中) TRADE_FINISHED(采购成功) TRADE_CLOSED(已关闭)
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetStatus() string {
    return r._status
}
// StartCreated Setter
// 起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetStartCreated(_startCreated string) error {
    r._startCreated = _startCreated
    r.Set("start_created", _startCreated)
    return nil
}

// StartCreated Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetStartCreated() string {
    return r._startCreated
}
// EndCreated Setter
// 结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetEndCreated(_endCreated string) error {
    r._endCreated = _endCreated
    r.Set("end_created", _endCreated)
    return nil
}

// EndCreated Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetEndCreated() string {
    return r._endCreated
}
// TimeType Setter
// 可选值：trade_time_type(采购单按照成交时间范围查询),update_time_type(采购单按照更新时间范围查询)
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetTimeType(_timeType string) error {
    r._timeType = _timeType
    r.Set("time_type", _timeType)
    return nil
}

// TimeType Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetTimeType() string {
    return r._timeType
}
// PageNo Setter
// 页码。（大于0的整数。默认为1）
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页条数。（每页条数不超过50条）
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// PurchaseOrderId Setter
// 采购单编号或分销流水号，若其它参数没传，则此参数必传。
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
    r._purchaseOrderId = _purchaseOrderId
    r.Set("purchase_order_id", _purchaseOrderId)
    return nil
}

// PurchaseOrderId Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetPurchaseOrderId() int64 {
    return r._purchaseOrderId
}
// Fields Setter
// 多个字段用","分隔。<br/><br/>fields<br/>如果为空：返回所有采购单对象(purchase_orders)字段。<br/>如果不为空：返回指定采购单对象(purchase_orders)字段。<br/><br/>例1：<br/>sub_purchase_orders.tc_order_id 表示只返回tc_order_id <br/>例2：<br/>sub_purchase_orders表示只返回子采购单列表
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetFields() string {
    return r._fields
}
// TcOrderId Setter
// 采购单下游买家订单id
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetTcOrderId(_tcOrderId int64) error {
    r._tcOrderId = _tcOrderId
    r.Set("tc_order_id", _tcOrderId)
    return nil
}

// TcOrderId Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetTcOrderId() int64 {
    return r._tcOrderId
}
// ChannelCodes Setter
// 渠道code，可批量 查询老供销：999
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetChannelCodes(_channelCodes []int64) error {
    r._channelCodes = _channelCodes
    r.Set("channel_codes", _channelCodes)
    return nil
}

// ChannelCodes Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetChannelCodes() []int64 {
    return r._channelCodes
}
// UserRoleType Setter
// 角色，供应商：2，分销商：1
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetUserRoleType(_userRoleType int64) error {
    r._userRoleType = _userRoleType
    r.Set("user_role_type", _userRoleType)
    return nil
}

// UserRoleType Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetUserRoleType() int64 {
    return r._userRoleType
}
// TradeTypes Setter
// 代销：1 经销：2 寄售（猫超自营寄售）：5 平台寄售：6
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetTradeTypes(_tradeTypes []int64) error {
    r._tradeTypes = _tradeTypes
    r.Set("trade_types", _tradeTypes)
    return nil
}

// TradeTypes Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetTradeTypes() []int64 {
    return r._tradeTypes
}
