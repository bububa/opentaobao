package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家的自定义区模板信息 API请求
cainiao.cloudprint.customares.get

供isv使用，获取商家的自定义区的模板信息
*/
type CainiaoCloudprintCustomaresGetRequest struct {
    model.Params
    // 用户使用的标准模板id
    _templateId   int64
}

// 初始化CainiaoCloudprintCustomaresGetRequest对象
func NewCainiaoCloudprintCustomaresGetRequest() *CainiaoCloudprintCustomaresGetRequest{
    return &CainiaoCloudprintCustomaresGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintCustomaresGetRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.customares.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintCustomaresGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateId Setter
// 用户使用的标准模板id
func (r *CainiaoCloudprintCustomaresGetRequest) SetTemplateId(_templateId int64) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r CainiaoCloudprintCustomaresGetRequest) GetTemplateId() int64 {
    return r._templateId
}
