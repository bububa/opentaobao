package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoRefundQueryAPIRequest 批量查询采购退款 API请求
// taobao.fenxiao.refund.query
//
// 供应商按查询条件批量查询代销采购退款
type TaobaoFenxiaoRefundQueryAPIRequest struct {
	model.Params
	// 查询的经营模式，当不指定时，默认查询代销订单 代销：1 经销：2 寄售（自营寄售）：5 平台寄售
	_tradeTypes []int64
	// 渠道市场编码，可批量指定。 当不指定时，按照配置的分销市场生效 渠道编码枚举：1-供销平台（淘宝）；2-供销平台（天猫）；3-供销平台（天猫超市）；5-供销平台（淘乡甜）；110001-供销平台（全球购）；110007-淘分销；200002-消费电子市场
	_channelCodes []int64
	// 查询退款单的修改时间开始,格式如:yyyy-MM-dd HH:mm:ss
	_startDate string
	// 查询退款单的修改时间结束,格式如:yyyy-MM-dd HH:mm:ss
	_endDate string
	// 页码（大于0的整数。无值或小于1的值按默认值1计）
	_pageNo int64
	// 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
	_pageSize int64
	// 角色，供应商：2，分销商：1
	_userRoleType int64
	// 退款流程类型：4：发货前退款；1：发货后退款不退货；2：发货后退款退货；3：售后退款；5：拒收；6：售后退货退款
	_refundFlowTypes int64
	// 是否查询下游消费者的退款信息
	_querySellerRefund bool
}

// NewTaobaoFenxiaoRefundQueryRequest 初始化TaobaoFenxiaoRefundQueryAPIRequest对象
func NewTaobaoFenxiaoRefundQueryRequest() *TaobaoFenxiaoRefundQueryAPIRequest {
	return &TaobaoFenxiaoRefundQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.refund.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTradeTypes is TradeTypes Setter
// 查询的经营模式，当不指定时，默认查询代销订单 代销：1 经销：2 寄售（自营寄售）：5 平台寄售
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetTradeTypes(_tradeTypes []int64) error {
	r._tradeTypes = _tradeTypes
	r.Set("trade_types", _tradeTypes)
	return nil
}

// GetTradeTypes TradeTypes Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetTradeTypes() []int64 {
	return r._tradeTypes
}

// SetChannelCodes is ChannelCodes Setter
// 渠道市场编码，可批量指定。 当不指定时，按照配置的分销市场生效 渠道编码枚举：1-供销平台（淘宝）；2-供销平台（天猫）；3-供销平台（天猫超市）；5-供销平台（淘乡甜）；110001-供销平台（全球购）；110007-淘分销；200002-消费电子市场
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetChannelCodes(_channelCodes []int64) error {
	r._channelCodes = _channelCodes
	r.Set("channel_codes", _channelCodes)
	return nil
}

// GetChannelCodes ChannelCodes Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetChannelCodes() []int64 {
	return r._channelCodes
}

// SetStartDate is StartDate Setter
// 查询退款单的修改时间开始,格式如:yyyy-MM-dd HH:mm:ss
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 查询退款单的修改时间结束,格式如:yyyy-MM-dd HH:mm:ss
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageNo is PageNo Setter
// 页码（大于0的整数。无值或小于1的值按默认值1计）
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetUserRoleType is UserRoleType Setter
// 角色，供应商：2，分销商：1
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetUserRoleType(_userRoleType int64) error {
	r._userRoleType = _userRoleType
	r.Set("user_role_type", _userRoleType)
	return nil
}

// GetUserRoleType UserRoleType Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetUserRoleType() int64 {
	return r._userRoleType
}

// SetRefundFlowTypes is RefundFlowTypes Setter
// 退款流程类型：4：发货前退款；1：发货后退款不退货；2：发货后退款退货；3：售后退款；5：拒收；6：售后退货退款
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetRefundFlowTypes(_refundFlowTypes int64) error {
	r._refundFlowTypes = _refundFlowTypes
	r.Set("refund_flow_types", _refundFlowTypes)
	return nil
}

// GetRefundFlowTypes RefundFlowTypes Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetRefundFlowTypes() int64 {
	return r._refundFlowTypes
}

// SetQuerySellerRefund is QuerySellerRefund Setter
// 是否查询下游消费者的退款信息
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetQuerySellerRefund(_querySellerRefund bool) error {
	r._querySellerRefund = _querySellerRefund
	r.Set("query_seller_refund", _querySellerRefund)
	return nil
}

// GetQuerySellerRefund QuerySellerRefund Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetQuerySellerRefund() bool {
	return r._querySellerRefund
}
