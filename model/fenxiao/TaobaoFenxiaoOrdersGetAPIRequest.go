package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoOrdersGetAPIRequest 查询采购单信息 API请求
// taobao.fenxiao.orders.get
//
// 查询代销采购单单据。
//
// 1. 支持商家按照供应商、分销商两种角色来查询数据。如果没有指定角色角色，系统会自动判断，此时如果商家存在供应商、分销商两种角色时，按照供应商角色查询。
// 2. 同时此接口还可以查询除供销经销外的其他经营模式的数据。如果需要查询供销经销单据请参考接口：taobao.fenxiao.dealer.requisitionorder.query
//
// 3. 发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。
type TaobaoFenxiaoOrdersGetAPIRequest struct {
	model.Params
	// 渠道市场编码，可批量指定。 当不指定时，按照配置的分销市场生效 渠道编码枚举：1-供销平台（淘宝）；2-供销平台（天猫）；3-供销平台（天猫超市）；5-供销平台（淘乡甜）；110001-供销平台（全球购）；110007-淘分销；200002-消费电子市场
	_channelCodes []int64
	// 查询的经营模式，当不指定时，默认查询代销订单 代销：1  经销：2  寄售（自营寄售）：5  平台寄售：6
	_tradeTypes []int64
	// 交易状态，不传默认查询所有采购单 根据用户角色选择自身状态可选值:  * 供应商： WAIT_SELLER_SEND_GOODS(等待发货)  WAIT_SELLER_CONFIRM_PAY(待确认收款)  WAIT_BUYER_PAY(等待付款)  WAIT_BUYER_CONFIRM_GOODS(已发货)  TRADE_REFUNDING(退款中)  TRADE_FINISHED(采购成功)  TRADE_CLOSED(已关闭)  * 分销商： WAIT_BUYER_PAY(等待付款)  WAIT_BUYER_CONFIRM_GOODS(待收货确认)   TRADE_FOR_PAY(已付款)  TRADE_REFUNDING(退款中)  TRADE_FINISHED(采购成功)  TRADE_CLOSED(已关闭)
	_status string
	// 起始时间，格式 yyyy-MM-dd HH:mm:ss 支持到秒的查询。若不传时分秒，默认为0时0分0秒。当指定了purchase_order_id或者tc_order_id时，此值可选，否则此参数必传。 结束时间和开始时间的时间间隔不能超过7天，精确到秒。
	_startCreated string
	// 结束时间，格式 yyyy-MM-dd HH:mm:ss 支持到秒的查询。若不传时分秒，默认为0时0分0秒。当指定了purchase_order_id或者tc_order_id时，此值可选，否则此参数必传。 结束时间和开始时间的时间间隔不能超过7天，精确到秒。
	_endCreated string
	// 时间类型： trade_time_type(默认类型，按照采购单创建时间范围查询，推荐按照此时间类型查询) update_time_type(采购单按照更新时间范围查询)
	_timeType string
	// 指定返回的字段（废弃该参数） 多个字段用","分隔。  fields 如果为空：返回所有采购单对象(purchase_orders)字段。 如果不为空：返回指定采购单对象(purchase_orders)字段。  例1： sub_purchase_orders.tc_order_id 表示只返回tc_order_id 例2： sub_purchase_orders表示只返回子采购单列表
	_fields string
	// 页码。（大于0的整数。默认为1）
	_pageNo int64
	// 每页条数。（每页条数不超过50条）
	_pageSize int64
	// 采购单编号或分销流水号。 当指定此参数后，其他可选参数可以为空
	_purchaseOrderId int64
	// 消费者交易单号（采购单下游买家订单id）。 当在代销、寄售交易中，此单号存在。
	_tcOrderId int64
	// 当前查询用户的角色 当不指定时，默认为供应商 供应商：2，分销商：1
	_userRoleType int64
}

// NewTaobaoFenxiaoOrdersGetRequest 初始化TaobaoFenxiaoOrdersGetAPIRequest对象
func NewTaobaoFenxiaoOrdersGetRequest() *TaobaoFenxiaoOrdersGetAPIRequest {
	return &TaobaoFenxiaoOrdersGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.orders.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetChannelCodes is ChannelCodes Setter
// 渠道市场编码，可批量指定。 当不指定时，按照配置的分销市场生效 渠道编码枚举：1-供销平台（淘宝）；2-供销平台（天猫）；3-供销平台（天猫超市）；5-供销平台（淘乡甜）；110001-供销平台（全球购）；110007-淘分销；200002-消费电子市场
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetChannelCodes(_channelCodes []int64) error {
	r._channelCodes = _channelCodes
	r.Set("channel_codes", _channelCodes)
	return nil
}

// GetChannelCodes ChannelCodes Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetChannelCodes() []int64 {
	return r._channelCodes
}

// SetTradeTypes is TradeTypes Setter
// 查询的经营模式，当不指定时，默认查询代销订单 代销：1  经销：2  寄售（自营寄售）：5  平台寄售：6
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetTradeTypes(_tradeTypes []int64) error {
	r._tradeTypes = _tradeTypes
	r.Set("trade_types", _tradeTypes)
	return nil
}

// GetTradeTypes TradeTypes Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetTradeTypes() []int64 {
	return r._tradeTypes
}

// SetStatus is Status Setter
// 交易状态，不传默认查询所有采购单 根据用户角色选择自身状态可选值:  * 供应商： WAIT_SELLER_SEND_GOODS(等待发货)  WAIT_SELLER_CONFIRM_PAY(待确认收款)  WAIT_BUYER_PAY(等待付款)  WAIT_BUYER_CONFIRM_GOODS(已发货)  TRADE_REFUNDING(退款中)  TRADE_FINISHED(采购成功)  TRADE_CLOSED(已关闭)  * 分销商： WAIT_BUYER_PAY(等待付款)  WAIT_BUYER_CONFIRM_GOODS(待收货确认)   TRADE_FOR_PAY(已付款)  TRADE_REFUNDING(退款中)  TRADE_FINISHED(采购成功)  TRADE_CLOSED(已关闭)
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetStatus() string {
	return r._status
}

// SetStartCreated is StartCreated Setter
// 起始时间，格式 yyyy-MM-dd HH:mm:ss 支持到秒的查询。若不传时分秒，默认为0时0分0秒。当指定了purchase_order_id或者tc_order_id时，此值可选，否则此参数必传。 结束时间和开始时间的时间间隔不能超过7天，精确到秒。
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetStartCreated(_startCreated string) error {
	r._startCreated = _startCreated
	r.Set("start_created", _startCreated)
	return nil
}

// GetStartCreated StartCreated Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetStartCreated() string {
	return r._startCreated
}

// SetEndCreated is EndCreated Setter
// 结束时间，格式 yyyy-MM-dd HH:mm:ss 支持到秒的查询。若不传时分秒，默认为0时0分0秒。当指定了purchase_order_id或者tc_order_id时，此值可选，否则此参数必传。 结束时间和开始时间的时间间隔不能超过7天，精确到秒。
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetEndCreated(_endCreated string) error {
	r._endCreated = _endCreated
	r.Set("end_created", _endCreated)
	return nil
}

// GetEndCreated EndCreated Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetEndCreated() string {
	return r._endCreated
}

// SetTimeType is TimeType Setter
// 时间类型： trade_time_type(默认类型，按照采购单创建时间范围查询，推荐按照此时间类型查询) update_time_type(采购单按照更新时间范围查询)
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetTimeType(_timeType string) error {
	r._timeType = _timeType
	r.Set("time_type", _timeType)
	return nil
}

// GetTimeType TimeType Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetTimeType() string {
	return r._timeType
}

// SetFields is Fields Setter
// 指定返回的字段（废弃该参数） 多个字段用&#34;,&#34;分隔。  fields 如果为空：返回所有采购单对象(purchase_orders)字段。 如果不为空：返回指定采购单对象(purchase_orders)字段。  例1： sub_purchase_orders.tc_order_id 表示只返回tc_order_id 例2： sub_purchase_orders表示只返回子采购单列表
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetFields() string {
	return r._fields
}

// SetPageNo is PageNo Setter
// 页码。（大于0的整数。默认为1）
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数。（每页条数不超过50条）
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPurchaseOrderId is PurchaseOrderId Setter
// 采购单编号或分销流水号。 当指定此参数后，其他可选参数可以为空
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
	r._purchaseOrderId = _purchaseOrderId
	r.Set("purchase_order_id", _purchaseOrderId)
	return nil
}

// GetPurchaseOrderId PurchaseOrderId Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetPurchaseOrderId() int64 {
	return r._purchaseOrderId
}

// SetTcOrderId is TcOrderId Setter
// 消费者交易单号（采购单下游买家订单id）。 当在代销、寄售交易中，此单号存在。
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetTcOrderId(_tcOrderId int64) error {
	r._tcOrderId = _tcOrderId
	r.Set("tc_order_id", _tcOrderId)
	return nil
}

// GetTcOrderId TcOrderId Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetTcOrderId() int64 {
	return r._tcOrderId
}

// SetUserRoleType is UserRoleType Setter
// 当前查询用户的角色 当不指定时，默认为供应商 供应商：2，分销商：1
func (r *TaobaoFenxiaoOrdersGetAPIRequest) SetUserRoleType(_userRoleType int64) error {
	r._userRoleType = _userRoleType
	r.Set("user_role_type", _userRoleType)
	return nil
}

// GetUserRoleType UserRoleType Getter
func (r TaobaoFenxiaoOrdersGetAPIRequest) GetUserRoleType() int64 {
	return r._userRoleType
}
