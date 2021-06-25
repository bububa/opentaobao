package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家使用的标准模板 APIRequest
cainiao.cloudprint.isvtemplates.get

获取商家使用的标准模板
*/
type CainiaoCloudprintIsvtemplatesGetRequest struct {
    model.Params

}

func NewCainiaoCloudprintIsvtemplatesGetRequest() *CainiaoCloudprintIsvtemplatesGetRequest{
    return &CainiaoCloudprintIsvtemplatesGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCloudprintIsvtemplatesGetRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.isvtemplates.get"
}

func (r CainiaoCloudprintIsvtemplatesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


