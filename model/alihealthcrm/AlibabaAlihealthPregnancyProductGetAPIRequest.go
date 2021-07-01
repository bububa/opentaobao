package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPregnancyProductGetAPIRequest
备孕首页获取达人配置商品 API请求
alibaba.alihealth.pregnancy.product.get

备孕首页获取达人配置商品 */
type AlibabaAlihealthPregnancyProductGetAPIRequest struct {
	model.Params
	// tab页对应id
	_sceneId int64
	// 起始页码，大于0
	_currentPage int64
	// 每页条数
	_pageSize int64
}

// NewAlibabaAlihealthPregnancyProductGetRequest 初始化AlibabaAlihealthPregnancyProductGetAPIRequest对象
func NewAlibabaAlihealthPregnancyProductGetRequest() *AlibabaAlihealthPregnancyProductGetAPIRequest {
	return &AlibabaAlihealthPregnancyProductGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyProductGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyProductGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SceneId Setter
// tab页对应id
func (r *AlibabaAlihealthPregnancyProductGetAPIRequest) SetSceneId(_sceneId int64) error {
	r._sceneId = _sceneId
	r.Set("scene_id", _sceneId)
	return nil
}

// Get SceneId Getter
func (r AlibabaAlihealthPregnancyProductGetAPIRequest) GetSceneId() int64 {
	return r._sceneId
}

// Set is CurrentPage Setter
// 起始页码，大于0
func (r *AlibabaAlihealthPregnancyProductGetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// Get CurrentPage Getter
func (r AlibabaAlihealthPregnancyProductGetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// Set is PageSize Setter
// 每页条数
func (r *AlibabaAlihealthPregnancyProductGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAlihealthPregnancyProductGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
