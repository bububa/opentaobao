package bill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaeBillsGetAPIRequest tae查询账单明细 API请求
// taobao.tae.bills.get
//
// tae查询账单明细
type TaobaoTaeBillsGetAPIRequest struct {
	model.Params
	// 传入需要返回的字段,参见Bill结构体
	_fields []string
	// 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
	_queryEndDate string
	// 开始时间
	_queryStartDate string
	// 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
	_queryDateType int64
	// 交易编号
	_pTradeId int64
	// 科目编号
	_itemId int64
	// 页数,建议不要超过100页,越大性能越低,有可能会超时
	_currentPage int64
	// 每页大小,默认40条,可选范围 ：40~100
	_pageSize int64
	// 子订单编号
	_tradeId int64
}

// NewTaobaoTaeBillsGetRequest 初始化TaobaoTaeBillsGetAPIRequest对象
func NewTaobaoTaeBillsGetRequest() *TaobaoTaeBillsGetAPIRequest {
	return &TaobaoTaeBillsGetAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTaeBillsGetAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r._queryEndDate = ""
	r._queryStartDate = ""
	r._queryDateType = 0
	r._pTradeId = 0
	r._itemId = 0
	r._currentPage = 0
	r._pageSize = 0
	r._tradeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaeBillsGetAPIRequest) GetApiMethodName() string {
	return "taobao.tae.bills.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaeBillsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaeBillsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 传入需要返回的字段,参见Bill结构体
func (r *TaobaoTaeBillsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTaeBillsGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetQueryEndDate is QueryEndDate Setter
// 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
func (r *TaobaoTaeBillsGetAPIRequest) SetQueryEndDate(_queryEndDate string) error {
	r._queryEndDate = _queryEndDate
	r.Set("query_end_date", _queryEndDate)
	return nil
}

// GetQueryEndDate QueryEndDate Getter
func (r TaobaoTaeBillsGetAPIRequest) GetQueryEndDate() string {
	return r._queryEndDate
}

// SetQueryStartDate is QueryStartDate Setter
// 开始时间
func (r *TaobaoTaeBillsGetAPIRequest) SetQueryStartDate(_queryStartDate string) error {
	r._queryStartDate = _queryStartDate
	r.Set("query_start_date", _queryStartDate)
	return nil
}

// GetQueryStartDate QueryStartDate Getter
func (r TaobaoTaeBillsGetAPIRequest) GetQueryStartDate() string {
	return r._queryStartDate
}

// SetQueryDateType is QueryDateType Setter
// 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
func (r *TaobaoTaeBillsGetAPIRequest) SetQueryDateType(_queryDateType int64) error {
	r._queryDateType = _queryDateType
	r.Set("query_date_type", _queryDateType)
	return nil
}

// GetQueryDateType QueryDateType Getter
func (r TaobaoTaeBillsGetAPIRequest) GetQueryDateType() int64 {
	return r._queryDateType
}

// SetPTradeId is PTradeId Setter
// 交易编号
func (r *TaobaoTaeBillsGetAPIRequest) SetPTradeId(_pTradeId int64) error {
	r._pTradeId = _pTradeId
	r.Set("p_trade_id", _pTradeId)
	return nil
}

// GetPTradeId PTradeId Getter
func (r TaobaoTaeBillsGetAPIRequest) GetPTradeId() int64 {
	return r._pTradeId
}

// SetItemId is ItemId Setter
// 科目编号
func (r *TaobaoTaeBillsGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoTaeBillsGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCurrentPage is CurrentPage Setter
// 页数,建议不要超过100页,越大性能越低,有可能会超时
func (r *TaobaoTaeBillsGetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoTaeBillsGetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页大小,默认40条,可选范围 ：40~100
func (r *TaobaoTaeBillsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTaeBillsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetTradeId is TradeId Setter
// 子订单编号
func (r *TaobaoTaeBillsGetAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaoTaeBillsGetAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

var poolTaobaoTaeBillsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTaeBillsGetRequest()
	},
}

// GetTaobaoTaeBillsGetRequest 从 sync.Pool 获取 TaobaoTaeBillsGetAPIRequest
func GetTaobaoTaeBillsGetAPIRequest() *TaobaoTaeBillsGetAPIRequest {
	return poolTaobaoTaeBillsGetAPIRequest.Get().(*TaobaoTaeBillsGetAPIRequest)
}

// ReleaseTaobaoTaeBillsGetAPIRequest 将 TaobaoTaeBillsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTaeBillsGetAPIRequest(v *TaobaoTaeBillsGetAPIRequest) {
	v.Reset()
	poolTaobaoTaeBillsGetAPIRequest.Put(v)
}
