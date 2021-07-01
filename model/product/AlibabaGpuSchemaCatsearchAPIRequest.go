package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGpuSchemaCatsearchAPIRequest
按类目查询spu接口 API请求
alibaba.gpu.schema.catsearch

按类目查询spu的schema接口 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.schema.catsearch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LeafCatId Setter
// 叶子类目ID
func (r *AlibabaGpuSchemaCatsearchAPIRequest) SetLeafCatId(_leafCatId int64) error {
	r._leafCatId = _leafCatId
	r.Set("leaf_cat_id", _leafCatId)
	return nil
}

// Get LeafCatId Getter
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetLeafCatId() int64 {
	return r._leafCatId
}

// Set is CurrentPage Setter
// 当前页
func (r *AlibabaGpuSchemaCatsearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// Get CurrentPage Getter
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// Set is PageSize Setter
// 每页大小
func (r *AlibabaGpuSchemaCatsearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is ProviderId Setter
// 渠道Id，如0代表天猫，8代表淘宝
func (r *AlibabaGpuSchemaCatsearchAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// Get ProviderId Getter
func (r AlibabaGpuSchemaCatsearchAPIRequest) GetProviderId() int64 {
	return r._providerId
}
