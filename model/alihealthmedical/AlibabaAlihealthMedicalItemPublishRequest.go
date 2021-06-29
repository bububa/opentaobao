package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方入驻-开通服务 APIRequest
alibaba.alihealth.medical.item.publish

三方入驻-开通服务
*/
type AlibabaAlihealthMedicalItemPublishRequest struct {
    model.Params

    // 请求
    request1   *ItemPublishRequest 

}

func NewAlibabaAlihealthMedicalItemPublishRequest() *AlibabaAlihealthMedicalItemPublishRequest{
    return &AlibabaAlihealthMedicalItemPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalItemPublishRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.item.publish"
}

func (r AlibabaAlihealthMedicalItemPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalItemPublishRequest) SetRequest1(request1 *ItemPublishRequest) error {
    r.request1 = request1
    r.Set("request1", request1)
    return nil
}

func (r AlibabaAlihealthMedicalItemPublishRequest) GetRequest1() *ItemPublishRequest {
    return r.request1
}

