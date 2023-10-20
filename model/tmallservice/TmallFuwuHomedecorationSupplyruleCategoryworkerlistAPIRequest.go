package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest 基于规则查品牌品类工人接口 API请求
// tmall.fuwu.homedecoration.supplyrule.categoryworkerlist
//
// 基于规则查品牌品类工人接口
type TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest struct {
	model.Params
	// 规则唯一编号
	_uniqueNo string
	// 当前页数
	_pageIndex int64
	// 每页大小，不传默认只查20条，最大不能超过500
	_pageSize int64
}

// NewTmallFuwuHomedecorationSupplyruleCategoryworkerlistRequest 初始化TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest对象
func NewTmallFuwuHomedecorationSupplyruleCategoryworkerlistRequest() *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest {
	return &TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) Reset() {
	r._uniqueNo = ""
	r._pageIndex = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) GetApiMethodName() string {
	return "tmall.fuwu.homedecoration.supplyrule.categoryworkerlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqueNo is UniqueNo Setter
// 规则唯一编号
func (r *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) SetUniqueNo(_uniqueNo string) error {
	r._uniqueNo = _uniqueNo
	r.Set("unique_no", _uniqueNo)
	return nil
}

// GetUniqueNo UniqueNo Getter
func (r TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) GetUniqueNo() string {
	return r._uniqueNo
}

// SetPageIndex is PageIndex Setter
// 当前页数
func (r *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页大小，不传默认只查20条，最大不能超过500
func (r *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallFuwuHomedecorationSupplyruleCategoryworkerlistRequest()
	},
}

// GetTmallFuwuHomedecorationSupplyruleCategoryworkerlistRequest 从 sync.Pool 获取 TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest
func GetTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest() *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest {
	return poolTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest.Get().(*TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest)
}

// ReleaseTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest 将 TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest 放入 sync.Pool
func ReleaseTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest(v *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest) {
	v.Reset()
	poolTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest.Put(v)
}
