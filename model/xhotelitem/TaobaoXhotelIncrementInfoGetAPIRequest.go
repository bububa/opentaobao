package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelIncrementInfoGetAPIRequest 酒店状态增量查询接口 API请求
// taobao.xhotel.increment.info.get
//
// 酒店状态增量查询接口
type TaobaoXhotelIncrementInfoGetAPIRequest struct {
	model.Params
	// 更新时间
	_gmtModified string
	// 分页参数：当前页数，从1开始计数。<br/>默认值：1
	_currentPage int64
	// 分页参数：每页酒店数量。最小值1，最大值为1000。默认值：1000
	_pageSize int64
	// 变化类别，1为酒店价格，0为酒店
	_changeType int64
}

// NewTaobaoXhotelIncrementInfoGetRequest 初始化TaobaoXhotelIncrementInfoGetAPIRequest对象
func NewTaobaoXhotelIncrementInfoGetRequest() *TaobaoXhotelIncrementInfoGetAPIRequest {
	return &TaobaoXhotelIncrementInfoGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelIncrementInfoGetAPIRequest) Reset() {
	r._gmtModified = ""
	r._currentPage = 0
	r._pageSize = 0
	r._changeType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelIncrementInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.increment.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelIncrementInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelIncrementInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGmtModified is GmtModified Setter
// 更新时间
func (r *TaobaoXhotelIncrementInfoGetAPIRequest) SetGmtModified(_gmtModified string) error {
	r._gmtModified = _gmtModified
	r.Set("gmt_modified", _gmtModified)
	return nil
}

// GetGmtModified GmtModified Getter
func (r TaobaoXhotelIncrementInfoGetAPIRequest) GetGmtModified() string {
	return r._gmtModified
}

// SetCurrentPage is CurrentPage Setter
// 分页参数：当前页数，从1开始计数。&lt;br/&gt;默认值：1
func (r *TaobaoXhotelIncrementInfoGetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoXhotelIncrementInfoGetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 分页参数：每页酒店数量。最小值1，最大值为1000。默认值：1000
func (r *TaobaoXhotelIncrementInfoGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoXhotelIncrementInfoGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetChangeType is ChangeType Setter
// 变化类别，1为酒店价格，0为酒店
func (r *TaobaoXhotelIncrementInfoGetAPIRequest) SetChangeType(_changeType int64) error {
	r._changeType = _changeType
	r.Set("change_type", _changeType)
	return nil
}

// GetChangeType ChangeType Getter
func (r TaobaoXhotelIncrementInfoGetAPIRequest) GetChangeType() int64 {
	return r._changeType
}

var poolTaobaoXhotelIncrementInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelIncrementInfoGetRequest()
	},
}

// GetTaobaoXhotelIncrementInfoGetRequest 从 sync.Pool 获取 TaobaoXhotelIncrementInfoGetAPIRequest
func GetTaobaoXhotelIncrementInfoGetAPIRequest() *TaobaoXhotelIncrementInfoGetAPIRequest {
	return poolTaobaoXhotelIncrementInfoGetAPIRequest.Get().(*TaobaoXhotelIncrementInfoGetAPIRequest)
}

// ReleaseTaobaoXhotelIncrementInfoGetAPIRequest 将 TaobaoXhotelIncrementInfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelIncrementInfoGetAPIRequest(v *TaobaoXhotelIncrementInfoGetAPIRequest) {
	v.Reset()
	poolTaobaoXhotelIncrementInfoGetAPIRequest.Put(v)
}
