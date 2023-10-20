package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderstatementgetAPIRequest 查询账单信息 API请求
// taobao.xhotel.order.statement.get
//
// 阿里根据此接口定义输出订单账务明细，结账状态发生变化时阿里需推送账单信息。系统商可实时调用该接口来查询订单的详情
type TaobaoxhotelorderstatementgetAPIRequest struct {
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

// NewTaobaoxhotelorderstatementgetRequest 初始化TaobaoxhotelorderstatementgetAPIRequest对象
func NewTaobaoxhotelorderstatementgetRequest() *TaobaoxhotelorderstatementgetAPIRequest {
	return &TaobaoxhotelorderstatementgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelorderstatementgetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.statement.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelorderstatementgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelorderstatementgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderTids is OrderTids Setter
// 要查询的tid列表，逗号分隔,列表查询;当此值不为空时候，其余参数忽略。最多单次20条。
func (r *TaobaoxhotelorderstatementgetAPIRequest) SetOrderTids(_orderTids string) error {
	r._orderTids = _orderTids
	r.Set("order_tids", _orderTids)
	return nil
}

// GetOrderTids OrderTids Getter
func (r TaobaoxhotelorderstatementgetAPIRequest) GetOrderTids() string {
	return r._orderTids
}

// SetTo is To Setter
// 查询结束时间
func (r *TaobaoxhotelorderstatementgetAPIRequest) SetTo(_to string) error {
	r._to = _to
	r.Set("to", _to)
	return nil
}

// GetTo To Getter
func (r TaobaoxhotelorderstatementgetAPIRequest) GetTo() string {
	return r._to
}

// SetFrom is From Setter
// 查询开始时间
func (r *TaobaoxhotelorderstatementgetAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaoxhotelorderstatementgetAPIRequest) GetFrom() string {
	return r._from
}

// SetHotelCode is HotelCode Setter
// 外部酒店编码
func (r *TaobaoxhotelorderstatementgetAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoxhotelorderstatementgetAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetVendor is Vendor Setter
// 系统商vendor
func (r *TaobaoxhotelorderstatementgetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelorderstatementgetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetPageSize is PageSize Setter
// 查询条数，最大支持500条
func (r *TaobaoxhotelorderstatementgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoxhotelorderstatementgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetStart is Start Setter
// 数据查询开始下标
func (r *TaobaoxhotelorderstatementgetAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TaobaoxhotelorderstatementgetAPIRequest) GetStart() int64 {
	return r._start
}

// SetDateType is DateType Setter
// 0：check_in, 1：check_out,2：分账时间
func (r *TaobaoxhotelorderstatementgetAPIRequest) SetDateType(_dateType int64) error {
	r._dateType = _dateType
	r.Set("date_type", _dateType)
	return nil
}

// GetDateType DateType Getter
func (r TaobaoxhotelorderstatementgetAPIRequest) GetDateType() int64 {
	return r._dateType
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoxhotelorderstatementgetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoxhotelorderstatementgetAPIRequest) GetTid() int64 {
	return r._tid
}
