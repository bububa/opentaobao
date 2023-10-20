package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuHomedecorationSupplyruleListAPIRequest 查询供给规则接口 API请求
// tmall.fuwu.homedecoration.supplyrule.list
//
// 查询供给规则接口
type TmallFuwuHomedecorationSupplyruleListAPIRequest struct {
	model.Params
	// 业务唯一编号
	_uniqueNo string
	// 每页大小，不传默认只查20条，最大不能超过500
	_pageSize int64
	// 当前页数
	_pageIndex int64
}

// NewTmallFuwuHomedecorationSupplyruleListRequest 初始化TmallFuwuHomedecorationSupplyruleListAPIRequest对象
func NewTmallFuwuHomedecorationSupplyruleListRequest() *TmallFuwuHomedecorationSupplyruleListAPIRequest {
	return &TmallFuwuHomedecorationSupplyruleListAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallFuwuHomedecorationSupplyruleListAPIRequest) Reset() {
	r._uniqueNo = ""
	r._pageSize = 0
	r._pageIndex = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFuwuHomedecorationSupplyruleListAPIRequest) GetApiMethodName() string {
	return "tmall.fuwu.homedecoration.supplyrule.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFuwuHomedecorationSupplyruleListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallFuwuHomedecorationSupplyruleListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqueNo is UniqueNo Setter
// 业务唯一编号
func (r *TmallFuwuHomedecorationSupplyruleListAPIRequest) SetUniqueNo(_uniqueNo string) error {
	r._uniqueNo = _uniqueNo
	r.Set("unique_no", _uniqueNo)
	return nil
}

// GetUniqueNo UniqueNo Getter
func (r TmallFuwuHomedecorationSupplyruleListAPIRequest) GetUniqueNo() string {
	return r._uniqueNo
}

// SetPageSize is PageSize Setter
// 每页大小，不传默认只查20条，最大不能超过500
func (r *TmallFuwuHomedecorationSupplyruleListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallFuwuHomedecorationSupplyruleListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 当前页数
func (r *TmallFuwuHomedecorationSupplyruleListAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TmallFuwuHomedecorationSupplyruleListAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

var poolTmallFuwuHomedecorationSupplyruleListAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallFuwuHomedecorationSupplyruleListRequest()
	},
}

// GetTmallFuwuHomedecorationSupplyruleListRequest 从 sync.Pool 获取 TmallFuwuHomedecorationSupplyruleListAPIRequest
func GetTmallFuwuHomedecorationSupplyruleListAPIRequest() *TmallFuwuHomedecorationSupplyruleListAPIRequest {
	return poolTmallFuwuHomedecorationSupplyruleListAPIRequest.Get().(*TmallFuwuHomedecorationSupplyruleListAPIRequest)
}

// ReleaseTmallFuwuHomedecorationSupplyruleListAPIRequest 将 TmallFuwuHomedecorationSupplyruleListAPIRequest 放入 sync.Pool
func ReleaseTmallFuwuHomedecorationSupplyruleListAPIRequest(v *TmallFuwuHomedecorationSupplyruleListAPIRequest) {
	v.Reset()
	poolTmallFuwuHomedecorationSupplyruleListAPIRequest.Put(v)
}
