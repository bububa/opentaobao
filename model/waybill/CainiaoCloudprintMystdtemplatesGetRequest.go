package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户使用的菜鸟电子面单模板信息 APIRequest
cainiao.cloudprint.mystdtemplates.get

获取用户使用的菜鸟电子面单
*/
type CainiaoCloudprintMystdtemplatesGetRequest struct {
    model.Params

}

func NewCainiaoCloudprintMystdtemplatesGetRequest() *CainiaoCloudprintMystdtemplatesGetRequest{
    return &CainiaoCloudprintMystdtemplatesGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCloudprintMystdtemplatesGetRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.mystdtemplates.get"
}

func (r CainiaoCloudprintMystdtemplatesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


