package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderStatementGetAPIRequest 查询账单信息 API请求
// taobao.xhotel.order.statement.get
//
// 阿里根据此接口定义输出订单账务明细，结账状态发生变化时阿里需推送账单信息。系统商可实时调用该接口来查询订单的详情
type TaobaoXhotelOrderStatementGetAPIRequest struct {
	model.Params
	// 要查询的tid列表，逗号分隔,列表查询;当此值不为空时候，其余参数忽略。最多单次20条。
	_orderTids string
	// 查询结束时间
	_to string
	// 查询开始时间
	_from string
	// 外部酒店编码
	_hotelCode string
	// 系统商vendor
	_vendor string
	// 查询条数，最大支持500条
	_pageSize int64
	// 数据查询开始下标
	_start int64
	// 0：check_in, 1：check_out,2：分账时间
	_dateType int64
	// 淘宝订单号
	_tid int64
}

// NewTaobaoXhotelOrderStatementGetRequest 初始化TaobaoXhotelOrderStatementGetAPIRequest对象
func NewTaobaoXhotelOrderStatementGetRequest() *TaobaoXhotelOrderStatementGetAPIRequest {
	return &TaobaoXhotelOrderStatementGetAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelOrderStatementGetAPIRequest) Reset() {
	r._orderTids = ""
	r._to = ""
	r._from = ""
	r._hotelCode = ""
	r._vendor = ""
	r._pageSize = 0
	r._start = 0
	r._dateType = 0
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.statement.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderTids is OrderTids Setter
// 要查询的tid列表，逗号分隔,列表查询;当此值不为空时候，其余参数忽略。最多单次20条。
func (r *TaobaoXhotelOrderStatementGetAPIRequest) SetOrderTids(_orderTids string) error {
	r._orderTids = _orderTids
	r.Set("order_tids", _orderTids)
	return nil
}

// GetOrderTids OrderTids Getter
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetOrderTids() string {
	return r._orderTids
}

// SetTo is To Setter
// 查询结束时间
func (r *TaobaoXhotelOrderStatementGetAPIRequest) SetTo(_to string) error {
	r._to = _to
	r.Set("to", _to)
	return nil
}

// GetTo To Getter
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetTo() string {
	return r._to
}

// SetFrom is From Setter
// 查询开始时间
func (r *TaobaoXhotelOrderStatementGetAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetFrom() string {
	return r._from
}

// SetHotelCode is HotelCode Setter
// 外部酒店编码
func (r *TaobaoXhotelOrderStatementGetAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetVendor is Vendor Setter
// 系统商vendor
func (r *TaobaoXhotelOrderStatementGetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetPageSize is PageSize Setter
// 查询条数，最大支持500条
func (r *TaobaoXhotelOrderStatementGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetStart is Start Setter
// 数据查询开始下标
func (r *TaobaoXhotelOrderStatementGetAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetStart() int64 {
	return r._start
}

// SetDateType is DateType Setter
// 0：check_in, 1：check_out,2：分账时间
func (r *TaobaoXhotelOrderStatementGetAPIRequest) SetDateType(_dateType int64) error {
	r._dateType = _dateType
	r.Set("date_type", _dateType)
	return nil
}

// GetDateType DateType Getter
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetDateType() int64 {
	return r._dateType
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoXhotelOrderStatementGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoXhotelOrderStatementGetAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoXhotelOrderStatementGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelOrderStatementGetRequest()
	},
}

// GetTaobaoXhotelOrderStatementGetRequest 从 sync.Pool 获取 TaobaoXhotelOrderStatementGetAPIRequest
func GetTaobaoXhotelOrderStatementGetAPIRequest() *TaobaoXhotelOrderStatementGetAPIRequest {
	return poolTaobaoXhotelOrderStatementGetAPIRequest.Get().(*TaobaoXhotelOrderStatementGetAPIRequest)
}

// ReleaseTaobaoXhotelOrderStatementGetAPIRequest 将 TaobaoXhotelOrderStatementGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelOrderStatementGetAPIRequest(v *TaobaoXhotelOrderStatementGetAPIRequest) {
	v.Reset()
	poolTaobaoXhotelOrderStatementGetAPIRequest.Put(v)
}
