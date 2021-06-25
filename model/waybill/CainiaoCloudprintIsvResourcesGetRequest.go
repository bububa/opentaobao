package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
isv资源查询 APIRequest
cainiao.cloudprint.isv.resources.get

isv资源查询，包括isv模板、打印项、预设的自定义区等
*/
type CainiaoCloudprintIsvResourcesGetRequest struct {
    model.Params

    // isv资源类型，分为：TEMPLATE（表示模板），PRINT_ITEM（打印项），CUSTOM_AREA（预设自定义区）
    isvResourceType   string 

}

func NewCainiaoCloudprintIsvResourcesGetRequest() *CainiaoCloudprintIsvResourcesGetRequest{
    return &CainiaoCloudprintIsvResourcesGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCloudprintIsvResourcesGetRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.isv.resources.get"
}

func (r CainiaoCloudprintIsvResourcesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCloudprintIsvResourcesGetRequest) SetIsvResourceType(isvResourceType string) error {
    r.isvResourceType = isvResourceType
    r.Set("isv_resource_type", isvResourceType)
    return nil
}

func (r CainiaoCloudprintIsvResourcesGetRequest) GetIsvResourceType() string {
    return r.isvResourceType
}

