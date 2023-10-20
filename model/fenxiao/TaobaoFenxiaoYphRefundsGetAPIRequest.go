package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoYphRefundsGetAPIRequest 一盘货商家批量查询退款单信息 API请求
// taobao.fenxiao.yph.refunds.get
//
// 一盘货商家批量查询退款单信息
type TaobaoFenxiaoYphRefundsGetAPIRequest struct {
	model.Params
	// 退款单最后修改时间，查询最后修改时间在modify_begin_time到modify_end_time时间间隔内的所有退款单，时间间隔不能超过30天，格式 yyyy-MM-dd HH:mm:ss
	_modifiedBeginTime string
	// 退款单最后修改时间，查询最后修改时间在modify_begin_time到modify_end_time时间间隔内的所有退款单，时间间隔不能超过30天，格式 yyyy-MM-dd HH:mm:ss
	_modifiedEndTime string
	// 退款单类型，10-未发货退款；20-已发货退货；30-已发货仅退款；40-拒收
	_refundTypeList int64
	// 交易模式（分销方式）：1-代销；2-经销；5-寄售；6-平台寄售
	_tradeTypes int64
	// 每页条数，必填，不可超过50条
	_pageSize int64
	// 页号，不填则默认为1
	_pageNum int64
	// 退款单状态，10-已撤回；20-等待卖家同意；30-卖家拒绝退款，等待买家修改；40-等待买家退货；50-买家退货，等待卖家确认收货；60-卖家拒绝收货；90-等待系统打款；100-退款成功；200-退款关闭
	_refundStatusList int64
	// 渠道市场编码，可批量指定。 当不指定时，按照配置的分销市场生效 渠道编码枚举：1-供销平台（淘宝）；2-供销平台（天猫）；3-供销平台（天猫超市）；5-供销平台（淘乡甜）；110001-供销平台（全球购）；110007-淘分销；200002-消费电子市场；302-猫超一盘货；506-虾选一盘货；27-优品一盘货
	_channelCodes int64
	// 分销采购单id
	_purchaseOrderId int64
	// 分销采购单子单id
	_subOrderId int64
	// 当前查询用户的角色 当不指定时，默认为供应商 供应商：2，分销商：1
	_userRoleType int64
	// 是否查询前台消费者退款
	_queryB2cRefund bool
}

// NewTaobaoFenxiaoYphRefundsGetRequest 初始化TaobaoFenxiaoYphRefundsGetAPIRequest对象
func NewTaobaoFenxiaoYphRefundsGetRequest() *TaobaoFenxiaoYphRefundsGetAPIRequest {
	return &TaobaoFenxiaoYphRefundsGetAPIRequest{
		Params: model.NewParams(12),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) Reset() {
	r._modifiedBeginTime = ""
	r._modifiedEndTime = ""
	r._refundTypeList = 0
	r._tradeTypes = 0
	r._pageSize = 0
	r._pageNum = 0
	r._refundStatusList = 0
	r._channelCodes = 0
	r._purchaseOrderId = 0
	r._subOrderId = 0
	r._userRoleType = 0
	r._queryB2cRefund = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.yph.refunds.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifiedBeginTime is ModifiedBeginTime Setter
// 退款单最后修改时间，查询最后修改时间在modify_begin_time到modify_end_time时间间隔内的所有退款单，时间间隔不能超过30天，格式 yyyy-MM-dd HH:mm:ss
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetModifiedBeginTime(_modifiedBeginTime string) error {
	r._modifiedBeginTime = _modifiedBeginTime
	r.Set("modified_begin_time", _modifiedBeginTime)
	return nil
}

// GetModifiedBeginTime ModifiedBeginTime Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetModifiedBeginTime() string {
	return r._modifiedBeginTime
}

// SetModifiedEndTime is ModifiedEndTime Setter
// 退款单最后修改时间，查询最后修改时间在modify_begin_time到modify_end_time时间间隔内的所有退款单，时间间隔不能超过30天，格式 yyyy-MM-dd HH:mm:ss
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetModifiedEndTime(_modifiedEndTime string) error {
	r._modifiedEndTime = _modifiedEndTime
	r.Set("modified_end_time", _modifiedEndTime)
	return nil
}

// GetModifiedEndTime ModifiedEndTime Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetModifiedEndTime() string {
	return r._modifiedEndTime
}

// SetRefundTypeList is RefundTypeList Setter
// 退款单类型，10-未发货退款；20-已发货退货；30-已发货仅退款；40-拒收
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetRefundTypeList(_refundTypeList int64) error {
	r._refundTypeList = _refundTypeList
	r.Set("refund_type_list", _refundTypeList)
	return nil
}

// GetRefundTypeList RefundTypeList Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetRefundTypeList() int64 {
	return r._refundTypeList
}

// SetTradeTypes is TradeTypes Setter
// 交易模式（分销方式）：1-代销；2-经销；5-寄售；6-平台寄售
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetTradeTypes(_tradeTypes int64) error {
	r._tradeTypes = _tradeTypes
	r.Set("trade_types", _tradeTypes)
	return nil
}

// GetTradeTypes TradeTypes Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetTradeTypes() int64 {
	return r._tradeTypes
}

// SetPageSize is PageSize Setter
// 每页条数，必填，不可超过50条
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 页号，不填则默认为1
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetRefundStatusList is RefundStatusList Setter
// 退款单状态，10-已撤回；20-等待卖家同意；30-卖家拒绝退款，等待买家修改；40-等待买家退货；50-买家退货，等待卖家确认收货；60-卖家拒绝收货；90-等待系统打款；100-退款成功；200-退款关闭
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetRefundStatusList(_refundStatusList int64) error {
	r._refundStatusList = _refundStatusList
	r.Set("refund_status_list", _refundStatusList)
	return nil
}

// GetRefundStatusList RefundStatusList Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetRefundStatusList() int64 {
	return r._refundStatusList
}

// SetChannelCodes is ChannelCodes Setter
// 渠道市场编码，可批量指定。 当不指定时，按照配置的分销市场生效 渠道编码枚举：1-供销平台（淘宝）；2-供销平台（天猫）；3-供销平台（天猫超市）；5-供销平台（淘乡甜）；110001-供销平台（全球购）；110007-淘分销；200002-消费电子市场；302-猫超一盘货；506-虾选一盘货；27-优品一盘货
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetChannelCodes(_channelCodes int64) error {
	r._channelCodes = _channelCodes
	r.Set("channel_codes", _channelCodes)
	return nil
}

// GetChannelCodes ChannelCodes Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetChannelCodes() int64 {
	return r._channelCodes
}

// SetPurchaseOrderId is PurchaseOrderId Setter
// 分销采购单id
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
	r._purchaseOrderId = _purchaseOrderId
	r.Set("purchase_order_id", _purchaseOrderId)
	return nil
}

// GetPurchaseOrderId PurchaseOrderId Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetPurchaseOrderId() int64 {
	return r._purchaseOrderId
}

// SetSubOrderId is SubOrderId Setter
// 分销采购单子单id
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}

// SetUserRoleType is UserRoleType Setter
// 当前查询用户的角色 当不指定时，默认为供应商 供应商：2，分销商：1
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetUserRoleType(_userRoleType int64) error {
	r._userRoleType = _userRoleType
	r.Set("user_role_type", _userRoleType)
	return nil
}

// GetUserRoleType UserRoleType Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetUserRoleType() int64 {
	return r._userRoleType
}

// SetQueryB2cRefund is QueryB2cRefund Setter
// 是否查询前台消费者退款
func (r *TaobaoFenxiaoYphRefundsGetAPIRequest) SetQueryB2cRefund(_queryB2cRefund bool) error {
	r._queryB2cRefund = _queryB2cRefund
	r.Set("query_b2c_refund", _queryB2cRefund)
	return nil
}

// GetQueryB2cRefund QueryB2cRefund Getter
func (r TaobaoFenxiaoYphRefundsGetAPIRequest) GetQueryB2cRefund() bool {
	return r._queryB2cRefund
}

var poolTaobaoFenxiaoYphRefundsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoYphRefundsGetRequest()
	},
}

// GetTaobaoFenxiaoYphRefundsGetRequest 从 sync.Pool 获取 TaobaoFenxiaoYphRefundsGetAPIRequest
func GetTaobaoFenxiaoYphRefundsGetAPIRequest() *TaobaoFenxiaoYphRefundsGetAPIRequest {
	return poolTaobaoFenxiaoYphRefundsGetAPIRequest.Get().(*TaobaoFenxiaoYphRefundsGetAPIRequest)
}

// ReleaseTaobaoFenxiaoYphRefundsGetAPIRequest 将 TaobaoFenxiaoYphRefundsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoYphRefundsGetAPIRequest(v *TaobaoFenxiaoYphRefundsGetAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoYphRefundsGetAPIRequest.Put(v)
}
