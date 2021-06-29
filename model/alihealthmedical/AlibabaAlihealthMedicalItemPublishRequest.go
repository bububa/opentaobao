package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方入驻-开通服务 API请求
alibaba.alihealth.medical.item.publish

三方入驻-开通服务
*/
type AlibabaAlihealthMedicalItemPublishRequest struct {
    model.Params
    // 请求
    request1   *ItemPublishRequest
}

// 初始化AlibabaAlihealthMedicalItemPublishRequest对象
func NewAlibabaAlihealthMedicalItemPublishRequest() *AlibabaAlihealthMedicalItemPublishRequest{
    return &AlibabaAlihealthMedicalItemPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalItemPublishRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.item.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalItemPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request1 Setter
// 请求
func (r *AlibabaAlihealthMedicalItemPublishRequest) SetRequest1(request1 *ItemPublishRequest) error {
    r.request1 = request1
    r.Set("request1", request1)
    return nil
}

// Request1 Getter
func (r AlibabaAlihealthMedicalItemPublishRequest) GetRequest1() *ItemPublishRequest {
    return r.request1
}
