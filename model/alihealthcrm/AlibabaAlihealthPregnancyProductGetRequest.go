package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
备孕首页获取达人配置商品 API请求
alibaba.alihealth.pregnancy.product.get

备孕首页获取达人配置商品
*/
type AlibabaAlihealthPregnancyProductGetRequest struct {
    model.Params
    // tab页对应id
    _sceneId   int64
    // 起始页码，大于0
    _currentPage   int64
    // 每页条数
    _pageSize   int64
}

// 初始化AlibabaAlihealthPregnancyProductGetRequest对象
func NewAlibabaAlihealthPregnancyProductGetRequest() *AlibabaAlihealthPregnancyProductGetRequest{
    return &AlibabaAlihealthPregnancyProductGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyProductGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.product.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyProductGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SceneId Setter
// tab页对应id
func (r *AlibabaAlihealthPregnancyProductGetRequest) SetSceneId(_sceneId int64) error {
    r._sceneId = _sceneId
    r.Set("scene_id", _sceneId)
    return nil
}

// SceneId Getter
func (r AlibabaAlihealthPregnancyProductGetRequest) GetSceneId() int64 {
    return r._sceneId
}
// CurrentPage Setter
// 起始页码，大于0
func (r *AlibabaAlihealthPregnancyProductGetRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaAlihealthPregnancyProductGetRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 每页条数
func (r *AlibabaAlihealthPregnancyProductGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthPregnancyProductGetRequest) GetPageSize() int64 {
    return r._pageSize
}
