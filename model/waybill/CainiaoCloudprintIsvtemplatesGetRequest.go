package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家使用的标准模板 API请求
cainiao.cloudprint.isvtemplates.get

获取商家使用的标准模板
*/
type CainiaoCloudprintIsvtemplatesGetRequest struct {
    model.Params
}

// 初始化CainiaoCloudprintIsvtemplatesGetRequest对象
func NewCainiaoCloudprintIsvtemplatesGetRequest() *CainiaoCloudprintIsvtemplatesGetRequest{
    return &CainiaoCloudprintIsvtemplatesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintIsvtemplatesGetRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.isvtemplates.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintIsvtemplatesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
