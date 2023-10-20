package bus

import (
	"net/url"
	"sync"

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
	// 结束时间 2017-04-23 13:33:43
	_toDate string
	// offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
	_type string
	// 分页大小 不超过1w
	_pageSize int64
	// 第几页 从1开始
	_pageIndex int64
}

// NewTaobaoBusHistoryorderGetRequest 初始化TaobaoBusHistoryorderGetAPIRequest对象
func NewTaobaoBusHistoryorderGetRequest() *TaobaoBusHistoryorderGetAPIRequest {
	return &TaobaoBusHistoryorderGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusHistoryorderGetAPIRequest) Reset() {
	r._fromDate = ""
	r._toDate = ""
	r._type = ""
	r._pageSize = 0
	r._pageIndex = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusHistoryorderGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.historyorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusHistoryorderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusHistoryorderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFromDate is FromDate Setter
// 开始时间 2017-04-23 13:33:43
func (r *TaobaoBusHistoryorderGetAPIRequest) SetFromDate(_fromDate string) error {
	r._fromDate = _fromDate
	r.Set("from_date", _fromDate)
	return nil
}

// GetFromDate FromDate Getter
func (r TaobaoBusHistoryorderGetAPIRequest) GetFromDate() string {
	return r._fromDate
}

// SetToDate is ToDate Setter
// 结束时间 2017-04-23 13:33:43
func (r *TaobaoBusHistoryorderGetAPIRequest) SetToDate(_toDate string) error {
	r._toDate = _toDate
	r.Set("to_date", _toDate)
	return nil
}

// GetToDate ToDate Getter
func (r TaobaoBusHistoryorderGetAPIRequest) GetToDate() string {
	return r._toDate
}

// SetType is Type Setter
// offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
func (r *TaobaoBusHistoryorderGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoBusHistoryorderGetAPIRequest) GetType() string {
	return r._type
}

// SetPageSize is PageSize Setter
// 分页大小 不超过1w
func (r *TaobaoBusHistoryorderGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoBusHistoryorderGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 第几页 从1开始
func (r *TaobaoBusHistoryorderGetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoBusHistoryorderGetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

var poolTaobaoBusHistoryorderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusHistoryorderGetRequest()
	},
}

// GetTaobaoBusHistoryorderGetRequest 从 sync.Pool 获取 TaobaoBusHistoryorderGetAPIRequest
func GetTaobaoBusHistoryorderGetAPIRequest() *TaobaoBusHistoryorderGetAPIRequest {
	return poolTaobaoBusHistoryorderGetAPIRequest.Get().(*TaobaoBusHistoryorderGetAPIRequest)
}

// ReleaseTaobaoBusHistoryorderGetAPIRequest 将 TaobaoBusHistoryorderGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusHistoryorderGetAPIRequest(v *TaobaoBusHistoryorderGetAPIRequest) {
	v.Reset()
	poolTaobaoBusHistoryorderGetAPIRequest.Put(v)
}
