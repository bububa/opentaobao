package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallfuwuhomedecorationsupplyrulelistAPIRequest 查询供给规则接口 API请求
// tmall.fuwu.homedecoration.supplyrule.list
//
// 查询供给规则接口
type TmallfuwuhomedecorationsupplyrulelistAPIRequest struct {
	model.Params
	// 业务唯一编号
	_uniqueNo string
	// 每页大小，不传默认只查20条，最大不能超过500
	_pageSize int64
	// 当前页数
	_pageIndex int64
}

// NewTmallfuwuhomedecorationsupplyrulelistRequest 初始化TmallfuwuhomedecorationsupplyrulelistAPIRequest对象
func NewTmallfuwuhomedecorationsupplyrulelistRequest() *TmallfuwuhomedecorationsupplyrulelistAPIRequest {
	return &TmallfuwuhomedecorationsupplyrulelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallfuwuhomedecorationsupplyrulelistAPIRequest) GetApiMethodName() string {
	return "tmall.fuwu.homedecoration.supplyrule.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallfuwuhomedecorationsupplyrulelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallfuwuhomedecorationsupplyrulelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqueNo is UniqueNo Setter
// 业务唯一编号
func (r *TmallfuwuhomedecorationsupplyrulelistAPIRequest) SetUniqueNo(_uniqueNo string) error {
	r._uniqueNo = _uniqueNo
	r.Set("unique_no", _uniqueNo)
	return nil
}

// GetUniqueNo UniqueNo Getter
func (r TmallfuwuhomedecorationsupplyrulelistAPIRequest) GetUniqueNo() string {
	return r._uniqueNo
}

// SetPageSize is PageSize Setter
// 每页大小，不传默认只查20条，最大不能超过500
func (r *TmallfuwuhomedecorationsupplyrulelistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallfuwuhomedecorationsupplyrulelistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 当前页数
func (r *TmallfuwuhomedecorationsupplyrulelistAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TmallfuwuhomedecorationsupplyrulelistAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
