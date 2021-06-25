package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取所有的菜鸟标准电子面单模板 APIRequest
cainiao.cloudprint.stdtemplates.get

获取菜鸟标准电子面单模板
*/
type CainiaoCloudprintStdtemplatesGetRequest struct {
    model.Params

}

func NewCainiaoCloudprintStdtemplatesGetRequest() *CainiaoCloudprintStdtemplatesGetRequest{
    return &CainiaoCloudprintStdtemplatesGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCloudprintStdtemplatesGetRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.stdtemplates.get"
}

func (r CainiaoCloudprintStdtemplatesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


