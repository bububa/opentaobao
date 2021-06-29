package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取所有的菜鸟标准电子面单模板 API请求
cainiao.cloudprint.stdtemplates.get

获取菜鸟标准电子面单模板
*/
type CainiaoCloudprintStdtemplatesGetRequest struct {
    model.Params
}

// 初始化CainiaoCloudprintStdtemplatesGetRequest对象
func NewCainiaoCloudprintStdtemplatesGetRequest() *CainiaoCloudprintStdtemplatesGetRequest{
    return &CainiaoCloudprintStdtemplatesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintStdtemplatesGetRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.stdtemplates.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintStdtemplatesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
