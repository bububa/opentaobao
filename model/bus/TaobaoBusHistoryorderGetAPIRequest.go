package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobushistoryordergetAPIRequest 历史订单查询（对账） API请求
// taobao.bus.historyorder.get
//
// 历史订单查询，对账接口
type TaobaobushistoryordergetAPIRequest struct {
	model.Params
	// 开始时间 2017-04-23 13:33:43
	_fromDate string
	// 结束时间 2017-04-23 13:33:43
	_toDate string
	// offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
	_type string
	// 分页大小 不超过1w
	_pageSize int64
	// 第几页 从1开始
	_pageIndex int64
}

// NewTaobaobushistoryordergetRequest 初始化TaobaobushistoryordergetAPIRequest对象
func NewTaobaobushistoryordergetRequest() *TaobaobushistoryordergetAPIRequest {
	return &TaobaobushistoryordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobushistoryordergetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.historyorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobushistoryordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobushistoryordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFromDate is FromDate Setter
// 开始时间 2017-04-23 13:33:43
func (r *TaobaobushistoryordergetAPIRequest) SetFromDate(_fromDate string) error {
	r._fromDate = _fromDate
	r.Set("from_date", _fromDate)
	return nil
}

// GetFromDate FromDate Getter
func (r TaobaobushistoryordergetAPIRequest) GetFromDate() string {
	return r._fromDate
}

// SetToDate is ToDate Setter
// 结束时间 2017-04-23 13:33:43
func (r *TaobaobushistoryordergetAPIRequest) SetToDate(_toDate string) error {
	r._toDate = _toDate
	r.Set("to_date", _toDate)
	return nil
}

// GetToDate ToDate Getter
func (r TaobaobushistoryordergetAPIRequest) GetToDate() string {
	return r._toDate
}

// SetType is Type Setter
// offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
func (r *TaobaobushistoryordergetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaobushistoryordergetAPIRequest) GetType() string {
	return r._type
}

// SetPageSize is PageSize Setter
// 分页大小 不超过1w
func (r *TaobaobushistoryordergetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaobushistoryordergetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 第几页 从1开始
func (r *TaobaobushistoryordergetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaobushistoryordergetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
