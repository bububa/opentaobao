package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGpuSchemaCatsearchAPIRequest 按类目查询spu接口 API请求
// alibaba.gpu.schema.catsearch
//
// 按类目查询spu的schema接口
type AlibabaGpuSchemaCatsearchAPIRequest struct {
	model.Params
	// 叶子类目ID
	_leafCatId int64
	// 当前页
	_currentPage int64
	// 每页大小
	_pageSize int64
	// 渠道Id，如0代表天猫，8代表淘宝
	_providerId int64
}

// NewAlibabaGpuSchemaCatsearchRequest 初始化AlibabaGpuSchemaCatsearchAPIRequest对象
func NewAlibabaGpuSchemaCatsearchRequest() *AlibabaGpuSchemaCatsearchAPIRequest {
	return &AlibabaGpuSchemaCatsearchAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaGpuSchemaCatsearchAPIRequest) Reset() {
	r._leafCatId = 0
	r._currentPage = 0
	r._pageSize = 0
	r._providerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.schema.catsearch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLeafCatId is LeafCatId Setter
// 叶子类目ID
func (r *AlibabaGpuSchemaCatsearchAPIRequest) SetLeafCatId(_leafCatId int64) error {
	r._leafCatId = _leafCatId
	r.Set("leaf_cat_id", _leafCatId)
	return nil
}

// GetLeafCatId LeafCatId Getter
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetLeafCatId() int64 {
	return r._leafCatId
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *AlibabaGpuSchemaCatsearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *AlibabaGpuSchemaCatsearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetProviderId is ProviderId Setter
// 渠道Id，如0代表天猫，8代表淘宝
func (r *AlibabaGpuSchemaCatsearchAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetProviderId() int64 {
	return r._providerId
}

var poolAlibabaGpuSchemaCatsearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaGpuSchemaCatsearchRequest()
	},
}

// GetAlibabaGpuSchemaCatsearchRequest 从 sync.Pool 获取 AlibabaGpuSchemaCatsearchAPIRequest
func GetAlibabaGpuSchemaCatsearchAPIRequest() *AlibabaGpuSchemaCatsearchAPIRequest {
	return poolAlibabaGpuSchemaCatsearchAPIRequest.Get().(*AlibabaGpuSchemaCatsearchAPIRequest)
}

// ReleaseAlibabaGpuSchemaCatsearchAPIRequest 将 AlibabaGpuSchemaCatsearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaGpuSchemaCatsearchAPIRequest(v *AlibabaGpuSchemaCatsearchAPIRequest) {
	v.Reset()
	poolAlibabaGpuSchemaCatsearchAPIRequest.Put(v)
}
