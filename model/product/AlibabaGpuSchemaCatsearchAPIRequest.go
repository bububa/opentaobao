package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabagpuschemacatsearchAPIRequest 按类目查询spu接口 API请求
// alibaba.gpu.schema.catsearch
//
// 按类目查询spu的schema接口
type AlibabagpuschemacatsearchAPIRequest struct {
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

// NewAlibabagpuschemacatsearchRequest 初始化AlibabagpuschemacatsearchAPIRequest对象
func NewAlibabagpuschemacatsearchRequest() *AlibabagpuschemacatsearchAPIRequest {
	return &AlibabagpuschemacatsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabagpuschemacatsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.schema.catsearch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabagpuschemacatsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabagpuschemacatsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLeafCatId is LeafCatId Setter
// 叶子类目ID
func (r *AlibabagpuschemacatsearchAPIRequest) SetLeafCatId(_leafCatId int64) error {
	r._leafCatId = _leafCatId
	r.Set("leaf_cat_id", _leafCatId)
	return nil
}

// GetLeafCatId LeafCatId Getter
func (r AlibabagpuschemacatsearchAPIRequest) GetLeafCatId() int64 {
	return r._leafCatId
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *AlibabagpuschemacatsearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabagpuschemacatsearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *AlibabagpuschemacatsearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabagpuschemacatsearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetProviderId is ProviderId Setter
// 渠道Id，如0代表天猫，8代表淘宝
func (r *AlibabagpuschemacatsearchAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabagpuschemacatsearchAPIRequest) GetProviderId() int64 {
	return r._providerId
}
