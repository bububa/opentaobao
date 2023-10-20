package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpregnancyproductgetAPIRequest 备孕首页获取达人配置商品 API请求
// alibaba.alihealth.pregnancy.product.get
//
// 备孕首页获取达人配置商品
type AlibabaalihealthpregnancyproductgetAPIRequest struct {
	model.Params
	// tab页对应id
	_sceneId int64
	// 起始页码，大于0
	_currentPage int64
	// 每页条数
	_pageSize int64
}

// NewAlibabaalihealthpregnancyproductgetRequest 初始化AlibabaalihealthpregnancyproductgetAPIRequest对象
func NewAlibabaalihealthpregnancyproductgetRequest() *AlibabaalihealthpregnancyproductgetAPIRequest {
	return &AlibabaalihealthpregnancyproductgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpregnancyproductgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpregnancyproductgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpregnancyproductgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSceneId is SceneId Setter
// tab页对应id
func (r *AlibabaalihealthpregnancyproductgetAPIRequest) SetSceneId(_sceneId int64) error {
	r._sceneId = _sceneId
	r.Set("scene_id", _sceneId)
	return nil
}

// GetSceneId SceneId Getter
func (r AlibabaalihealthpregnancyproductgetAPIRequest) GetSceneId() int64 {
	return r._sceneId
}

// SetCurrentPage is CurrentPage Setter
// 起始页码，大于0
func (r *AlibabaalihealthpregnancyproductgetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaalihealthpregnancyproductgetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *AlibabaalihealthpregnancyproductgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthpregnancyproductgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
