package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobillbillsgetAPIRequest 查询账单明细数据(自研发商家专用) API请求
// taobao.bill.bills.get
//
// 查询账单明细数据
type TaobaobillbillsgetAPIRequest struct {
	model.Params
	// 传入需要返回的字段,参见Bill结构体
	_fields []string
	// 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
	_endTime string
	// 开始时间
	_startTime string
	// 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
	_timeType int64
	// 交易编号
	_tradeId int64
	// 科目编号
	_accountId int64
	// 页数,建议不要超过100页,越大性能越低,有可能会超时
	_pageNo int64
	// 每页大小,默认40条,可选范围 ：40~100
	_pageSize int64
	// 子订单编号
	_orderId int64
}

// NewTaobaobillbillsgetRequest 初始化TaobaobillbillsgetAPIRequest对象
func NewTaobaobillbillsgetRequest() *TaobaobillbillsgetAPIRequest {
	return &TaobaobillbillsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobillbillsgetAPIRequest) GetApiMethodName() string {
	return "taobao.bill.bills.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobillbillsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobillbillsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 传入需要返回的字段,参见Bill结构体
func (r *TaobaobillbillsgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaobillbillsgetAPIRequest) GetFields() []string {
	return r._fields
}

// SetEndTime is EndTime Setter
// 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
func (r *TaobaobillbillsgetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaobillbillsgetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaobillbillsgetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaobillbillsgetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetTimeType is TimeType Setter
// 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
func (r *TaobaobillbillsgetAPIRequest) SetTimeType(_timeType int64) error {
	r._timeType = _timeType
	r.Set("time_type", _timeType)
	return nil
}

// GetTimeType TimeType Getter
func (r TaobaobillbillsgetAPIRequest) GetTimeType() int64 {
	return r._timeType
}

// SetTradeId is TradeId Setter
// 交易编号
func (r *TaobaobillbillsgetAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaobillbillsgetAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

// SetAccountId is AccountId Setter
// 科目编号
func (r *TaobaobillbillsgetAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r TaobaobillbillsgetAPIRequest) GetAccountId() int64 {
	return r._accountId
}

// SetPageNo is PageNo Setter
// 页数,建议不要超过100页,越大性能越低,有可能会超时
func (r *TaobaobillbillsgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaobillbillsgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小,默认40条,可选范围 ：40~100
func (r *TaobaobillbillsgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaobillbillsgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetOrderId is OrderId Setter
// 子订单编号
func (r *TaobaobillbillsgetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaobillbillsgetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
