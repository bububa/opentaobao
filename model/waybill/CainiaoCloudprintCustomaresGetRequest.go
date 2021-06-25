package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家的自定义区模板信息 APIRequest
cainiao.cloudprint.customares.get

供isv使用，获取商家的自定义区的模板信息
*/
type CainiaoCloudprintCustomaresGetRequest struct {
    model.Params

    // 用户使用的标准模板id
    templateId   int64 

}

func NewCainiaoCloudprintCustomaresGetRequest() *CainiaoCloudprintCustomaresGetRequest{
    return &CainiaoCloudprintCustomaresGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCloudprintCustomaresGetRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.customares.get"
}

func (r CainiaoCloudprintCustomaresGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCloudprintCustomaresGetRequest) SetTemplateId(templateId int64) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r CainiaoCloudprintCustomaresGetRequest) GetTemplateId() int64 {
    return r.templateId
}

