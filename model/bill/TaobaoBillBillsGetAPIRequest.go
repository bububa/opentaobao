package bill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBillBillsGetAPIRequest 查询账单明细数据(自研发商家专用) API请求
// taobao.bill.bills.get
//
// 查询账单明细数据
type TaobaoBillBillsGetAPIRequest struct {
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

// NewTaobaoBillBillsGetRequest 初始化TaobaoBillBillsGetAPIRequest对象
func NewTaobaoBillBillsGetRequest() *TaobaoBillBillsGetAPIRequest {
	return &TaobaoBillBillsGetAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBillBillsGetAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r._endTime = ""
	r._startTime = ""
	r._timeType = 0
	r._tradeId = 0
	r._accountId = 0
	r._pageNo = 0
	r._pageSize = 0
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBillBillsGetAPIRequest) GetApiMethodName() string {
	return "taobao.bill.bills.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBillBillsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBillBillsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 传入需要返回的字段,参见Bill结构体
func (r *TaobaoBillBillsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoBillBillsGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetEndTime is EndTime Setter
// 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
func (r *TaobaoBillBillsGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoBillBillsGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoBillBillsGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoBillBillsGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetTimeType is TimeType Setter
// 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
func (r *TaobaoBillBillsGetAPIRequest) SetTimeType(_timeType int64) error {
	r._timeType = _timeType
	r.Set("time_type", _timeType)
	return nil
}

// GetTimeType TimeType Getter
func (r TaobaoBillBillsGetAPIRequest) GetTimeType() int64 {
	return r._timeType
}

// SetTradeId is TradeId Setter
// 交易编号
func (r *TaobaoBillBillsGetAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaoBillBillsGetAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

// SetAccountId is AccountId Setter
// 科目编号
func (r *TaobaoBillBillsGetAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r TaobaoBillBillsGetAPIRequest) GetAccountId() int64 {
	return r._accountId
}

// SetPageNo is PageNo Setter
// 页数,建议不要超过100页,越大性能越低,有可能会超时
func (r *TaobaoBillBillsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoBillBillsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小,默认40条,可选范围 ：40~100
func (r *TaobaoBillBillsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoBillBillsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetOrderId is OrderId Setter
// 子订单编号
func (r *TaobaoBillBillsGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoBillBillsGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTaobaoBillBillsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBillBillsGetRequest()
	},
}

// GetTaobaoBillBillsGetRequest 从 sync.Pool 获取 TaobaoBillBillsGetAPIRequest
func GetTaobaoBillBillsGetAPIRequest() *TaobaoBillBillsGetAPIRequest {
	return poolTaobaoBillBillsGetAPIRequest.Get().(*TaobaoBillBillsGetAPIRequest)
}

// ReleaseTaobaoBillBillsGetAPIRequest 将 TaobaoBillBillsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBillBillsGetAPIRequest(v *TaobaoBillBillsGetAPIRequest) {
	v.Reset()
	poolTaobaoBillBillsGetAPIRequest.Put(v)
}
