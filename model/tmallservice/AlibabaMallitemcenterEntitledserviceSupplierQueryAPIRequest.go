package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest 根据天猫id查询门店服务授权 API请求
// alibaba.mallitemcenter.entitledservice.supplier.query
//
// 根据天猫id查询门店服务授权
type AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest struct {
	model.Params
	// 天猫id
	_id int64
	// 第几页
	_currentPage int64
	// 每页条数
	_pageSize int64
}

// NewAlibabaMallitemcenterEntitledserviceSupplierQueryRequest 初始化AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest对象
func NewAlibabaMallitemcenterEntitledserviceSupplierQueryRequest() *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest {
	return &AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) Reset() {
	r._id = 0
	r._currentPage = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mallitemcenter.entitledservice.supplier.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 天猫id
func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetId() int64 {
	return r._id
}

// SetCurrentPage is CurrentPage Setter
// 第几页
func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMallitemcenterEntitledserviceSupplierQueryRequest()
	},
}

// GetAlibabaMallitemcenterEntitledserviceSupplierQueryRequest 从 sync.Pool 获取 AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest
func GetAlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest() *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest {
	return poolAlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest.Get().(*AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest)
}

// ReleaseAlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest 将 AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest(v *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) {
	v.Reset()
	poolAlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest.Put(v)
}
