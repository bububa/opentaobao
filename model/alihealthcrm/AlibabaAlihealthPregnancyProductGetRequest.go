package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
备孕首页获取达人配置商品 APIRequest
alibaba.alihealth.pregnancy.product.get

备孕首页获取达人配置商品
*/
type AlibabaAlihealthPregnancyProductGetRequest struct {
    model.Params

    // tab页对应id
    sceneId   int64 

    // 起始页码，大于0
    currentPage   int64 

    // 每页条数
    pageSize   int64 

}

func NewAlibabaAlihealthPregnancyProductGetRequest() *AlibabaAlihealthPregnancyProductGetRequest{
    return &AlibabaAlihealthPregnancyProductGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthPregnancyProductGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.product.get"
}

func (r AlibabaAlihealthPregnancyProductGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthPregnancyProductGetRequest) SetSceneId(sceneId int64) error {
    r.sceneId = sceneId
    r.Set("scene_id", sceneId)
    return nil
}

func (r AlibabaAlihealthPregnancyProductGetRequest) GetSceneId() int64 {
    return r.sceneId
}

func (r *AlibabaAlihealthPregnancyProductGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlibabaAlihealthPregnancyProductGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlibabaAlihealthPregnancyProductGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthPregnancyProductGetRequest) GetPageSize() int64 {
    return r.pageSize
}

