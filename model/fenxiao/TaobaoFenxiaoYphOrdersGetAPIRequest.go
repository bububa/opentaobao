package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoYphOrdersGetAPIRequest 批量查询一盘货采购单信息 API请求
// taobao.fenxiao.yph.orders.get
//
// 一盘货商家批量查询采购单信息
type TaobaoFenxiaoYphOrdersGetAPIRequest struct {
	model.Params
	// 采购单最后更新时间-起始时间，格式 yyyy-MM-dd HH:mm:ss 支持到秒的查询。时间跨度必须在0-7天。
	_beginTime string
	// 采购单最后更新时间-结束时间，格式 yyyy-MM-dd HH:mm:ss 支持到秒的查询。时间跨度必须在0-7天。
	_endTime string
	// 查询的经营模式，当不指定时，默认查询代销订单 代销：1 经销：2 寄售（自营寄售）：5 平台寄售：6
	_tradeTypes int64
	// 每页条数。（每页条数不超过50条）
	_pageSize int64
	// 页号，不填则默认为1
	_pageNum int64
	// 当前查询用户的角色 当不指定时，默认为供应商 供应商：2，分销商：1
	_userRoleType int64
	// 渠道市场编码，可批量指定。 当不指定时，按照配置的分销市场生效 渠道编码枚举：1-供销平台（淘宝）；2-供销平台（天猫）；3-供销平台（天猫超市）；5-供销平台（淘乡甜）；110001-供销平台（全球购）；110007-淘分销；200002-消费电子市场；302-猫超一盘货；506-虾选一盘货
	_channelCodes int64
	// 主采购单交易状态-数字表示，枚举：1-已付款，待发货；2-待确认收款；4-等待买家付款；5-已付款，已发货；6-交易成功；7-交易关闭；8-已付款；
	_orderStatus int64
	// 主采购单退款状态-数字表示，枚举：9-没有过申请退款；10-退款活动中,至少有一个子单在退款中；11-退款结束，即曾经发生过退款且所有退款都完结了
	_refundStatus int64
}

// NewTaobaoFenxiaoYphOrdersGetRequest 初始化TaobaoFenxiaoYphOrdersGetAPIRequest对象
func NewTaobaoFenxiaoYphOrdersGetRequest() *TaobaoFenxiaoYphOrdersGetAPIRequest {
	return &TaobaoFenxiaoYphOrdersGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.yph.orders.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBeginTime is BeginTime Setter
// 采购单最后更新时间-起始时间，格式 yyyy-MM-dd HH:mm:ss 支持到秒的查询。时间跨度必须在0-7天。
func (r *TaobaoFenxiaoYphOrdersGetAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 采购单最后更新时间-结束时间，格式 yyyy-MM-dd HH:mm:ss 支持到秒的查询。时间跨度必须在0-7天。
func (r *TaobaoFenxiaoYphOrdersGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetTradeTypes is TradeTypes Setter
// 查询的经营模式，当不指定时，默认查询代销订单 代销：1 经销：2 寄售（自营寄售）：5 平台寄售：6
func (r *TaobaoFenxiaoYphOrdersGetAPIRequest) SetTradeTypes(_tradeTypes int64) error {
	r._tradeTypes = _tradeTypes
	r.Set("trade_types", _tradeTypes)
	return nil
}

// GetTradeTypes TradeTypes Getter
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetTradeTypes() int64 {
	return r._tradeTypes
}

// SetPageSize is PageSize Setter
// 每页条数。（每页条数不超过50条）
func (r *TaobaoFenxiaoYphOrdersGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 页号，不填则默认为1
func (r *TaobaoFenxiaoYphOrdersGetAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetUserRoleType is UserRoleType Setter
// 当前查询用户的角色 当不指定时，默认为供应商 供应商：2，分销商：1
func (r *TaobaoFenxiaoYphOrdersGetAPIRequest) SetUserRoleType(_userRoleType int64) error {
	r._userRoleType = _userRoleType
	r.Set("user_role_type", _userRoleType)
	return nil
}

// GetUserRoleType UserRoleType Getter
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetUserRoleType() int64 {
	return r._userRoleType
}

// SetChannelCodes is ChannelCodes Setter
// 渠道市场编码，可批量指定。 当不指定时，按照配置的分销市场生效 渠道编码枚举：1-供销平台（淘宝）；2-供销平台（天猫）；3-供销平台（天猫超市）；5-供销平台（淘乡甜）；110001-供销平台（全球购）；110007-淘分销；200002-消费电子市场；302-猫超一盘货；506-虾选一盘货
func (r *TaobaoFenxiaoYphOrdersGetAPIRequest) SetChannelCodes(_channelCodes int64) error {
	r._channelCodes = _channelCodes
	r.Set("channel_codes", _channelCodes)
	return nil
}

// GetChannelCodes ChannelCodes Getter
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetChannelCodes() int64 {
	return r._channelCodes
}

// SetOrderStatus is OrderStatus Setter
// 主采购单交易状态-数字表示，枚举：1-已付款，待发货；2-待确认收款；4-等待买家付款；5-已付款，已发货；6-交易成功；7-交易关闭；8-已付款；
func (r *TaobaoFenxiaoYphOrdersGetAPIRequest) SetOrderStatus(_orderStatus int64) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetOrderStatus() int64 {
	return r._orderStatus
}

// SetRefundStatus is RefundStatus Setter
// 主采购单退款状态-数字表示，枚举：9-没有过申请退款；10-退款活动中,至少有一个子单在退款中；11-退款结束，即曾经发生过退款且所有退款都完结了
func (r *TaobaoFenxiaoYphOrdersGetAPIRequest) SetRefundStatus(_refundStatus int64) error {
	r._refundStatus = _refundStatus
	r.Set("refund_status", _refundStatus)
	return nil
}

// GetRefundStatus RefundStatus Getter
func (r TaobaoFenxiaoYphOrdersGetAPIRequest) GetRefundStatus() int64 {
	return r._refundStatus
}
