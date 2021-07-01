package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户使用的菜鸟电子面单模板信息 API请求
cainiao.cloudprint.mystdtemplates.get

获取用户使用的菜鸟电子面单
*/
type CainiaoCloudprintMystdtemplatesGetAPIRequest struct {
    model.Params
}

// 初始化CainiaoCloudprintMystdtemplatesGetAPIRequest对象
func NewCainiaoCloudprintMystdtemplatesGetRequest() *CainiaoCloudprintMystdtemplatesGetAPIRequest{
    return &CainiaoCloudprintMystdtemplatesGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintMystdtemplatesGetAPIRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.mystdtemplates.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintMystdtemplatesGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
