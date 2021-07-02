package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusHistoryorderGetAPIRequest 历史订单查询（对账） API请求
// taobao.bus.historyorder.get
//
// 历史订单查询，对账接口
type TaobaoBusHistoryorderGetAPIRequest struct {
	model.Params
	// 开始时间 2017-04-23 13:33:43
	_fromDate string
	// 分页大小 不超过1w
	_pageSize int64
	// 结束时间 2017-04-23 13:33:43
	_toDate string
	// offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
	_type string
	// 第几页 从1开始
	_pageIndex int64
}

// NewTaobaoBusHistoryorderGetRequest 初始化TaobaoBusHistoryorderGetAPIRequest对象
func NewTaobaoBusHistoryorderGetRequest() *TaobaoBusHistoryorderGetAPIRequest {
	return &TaobaoBusHistoryorderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusHistoryorderGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.historyorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusHistoryorderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FromDate Setter
// 开始时间 2017-04-23 13:33:43
func (r *TaobaoBusHistoryorderGetAPIRequest) SetFromDate(_fromDate string) error {
	r._fromDate = _fromDate
	r.Set("from_date", _fromDate)
	return nil
}

// Get FromDate Getter
func (r TaobaoBusHistoryorderGetAPIRequest) GetFromDate() string {
	return r._fromDate
}

// Set is PageSize Setter
// 分页大小 不超过1w
func (r *TaobaoBusHistoryorderGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoBusHistoryorderGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is ToDate Setter
// 结束时间 2017-04-23 13:33:43
func (r *TaobaoBusHistoryorderGetAPIRequest) SetToDate(_toDate string) error {
	r._toDate = _toDate
	r.Set("to_date", _toDate)
	return nil
}

// Get ToDate Getter
func (r TaobaoBusHistoryorderGetAPIRequest) GetToDate() string {
	return r._toDate
}

// Set is Type Setter
// offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
func (r *TaobaoBusHistoryorderGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoBusHistoryorderGetAPIRequest) GetType() string {
	return r._type
}

// Set is PageIndex Setter
// 第几页 从1开始
func (r *TaobaoBusHistoryorderGetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// Get PageIndex Getter
func (r TaobaoBusHistoryorderGetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
