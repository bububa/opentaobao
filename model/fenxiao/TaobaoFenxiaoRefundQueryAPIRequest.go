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
	// 代销采购退款单最早修改时间
	_startDate string
	// 代销采购退款最迟修改时间。与start_date的最大时间间隔不能超过30天
	_endDate string
	// 页码（大于0的整数。无值或小于1的值按默认值1计）
	_pageNo int64
	// 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
	_pageSize int64
	// 是否查询下游买家的退款信息
	_querySellerRefund bool
	// 渠道code，可批量 老供销渠道：999
	_tradeTypes []int64
	// 角色，供应商：2，分销商：1
	_userRoleType int64
	// 代销：1 经销：2 寄售（猫超自营寄售）：5 平台寄售：6
	_channelCodes []int64
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

// Set is StartDate Setter
// 代销采购退款单最早修改时间
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is EndDate Setter
// 代销采购退款最迟修改时间。与start_date的最大时间间隔不能超过30天
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is PageNo Setter
// 页码（大于0的整数。无值或小于1的值按默认值1计）
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is QuerySellerRefund Setter
// 是否查询下游买家的退款信息
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetQuerySellerRefund(_querySellerRefund bool) error {
	r._querySellerRefund = _querySellerRefund
	r.Set("query_seller_refund", _querySellerRefund)
	return nil
}

// Get QuerySellerRefund Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetQuerySellerRefund() bool {
	return r._querySellerRefund
}

// Set is TradeTypes Setter
// 渠道code，可批量 老供销渠道：999
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetTradeTypes(_tradeTypes []int64) error {
	r._tradeTypes = _tradeTypes
	r.Set("trade_types", _tradeTypes)
	return nil
}

// Get TradeTypes Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetTradeTypes() []int64 {
	return r._tradeTypes
}

// Set is UserRoleType Setter
// 角色，供应商：2，分销商：1
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetUserRoleType(_userRoleType int64) error {
	r._userRoleType = _userRoleType
	r.Set("user_role_type", _userRoleType)
	return nil
}

// Get UserRoleType Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetUserRoleType() int64 {
	return r._userRoleType
}

// Set is ChannelCodes Setter
// 代销：1 经销：2 寄售（猫超自营寄售）：5 平台寄售：6
func (r *TaobaoFenxiaoRefundQueryAPIRequest) SetChannelCodes(_channelCodes []int64) error {
	r._channelCodes = _channelCodes
	r.Set("channel_codes", _channelCodes)
	return nil
}

// Get ChannelCodes Getter
func (r TaobaoFenxiaoRefundQueryAPIRequest) GetChannelCodes() []int64 {
	return r._channelCodes
}
