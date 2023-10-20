package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamallitemcenterentitledservicesupplierqueryAPIRequest 根据天猫id查询门店服务授权 API请求
// alibaba.mallitemcenter.entitledservice.supplier.query
//
// 根据天猫id查询门店服务授权
type AlibabamallitemcenterentitledservicesupplierqueryAPIRequest struct {
	model.Params
	// 天猫id
	_id int64
	// 第几页
	_currentPage int64
	// 每页条数
	_pageSize int64
}

// NewAlibabamallitemcenterentitledservicesupplierqueryRequest 初始化AlibabamallitemcenterentitledservicesupplierqueryAPIRequest对象
func NewAlibabamallitemcenterentitledservicesupplierqueryRequest() *AlibabamallitemcenterentitledservicesupplierqueryAPIRequest {
	return &AlibabamallitemcenterentitledservicesupplierqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamallitemcenterentitledservicesupplierqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mallitemcenter.entitledservice.supplier.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamallitemcenterentitledservicesupplierqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamallitemcenterentitledservicesupplierqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 天猫id
func (r *AlibabamallitemcenterentitledservicesupplierqueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabamallitemcenterentitledservicesupplierqueryAPIRequest) GetId() int64 {
	return r._id
}

// SetCurrentPage is CurrentPage Setter
// 第几页
func (r *AlibabamallitemcenterentitledservicesupplierqueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabamallitemcenterentitledservicesupplierqueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *AlibabamallitemcenterentitledservicesupplierqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabamallitemcenterentitledservicesupplierqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
