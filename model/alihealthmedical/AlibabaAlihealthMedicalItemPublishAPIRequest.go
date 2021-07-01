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
type AlibabaAlihealthMedicalItemPublishAPIRequest struct {
    model.Params
    // 请求
    _request1   *ItemPublishRequest
}

// 初始化AlibabaAlihealthMedicalItemPublishAPIRequest对象
func NewAlibabaAlihealthMedicalItemPublishRequest() *AlibabaAlihealthMedicalItemPublishAPIRequest{
    return &AlibabaAlihealthMedicalItemPublishAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalItemPublishAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.item.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalItemPublishAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request1 Setter
// 请求
func (r *AlibabaAlihealthMedicalItemPublishAPIRequest) SetRequest1(_request1 *ItemPublishRequest) error {
    r._request1 = _request1
    r.Set("request1", _request1)
    return nil
}

// Request1 Getter
func (r AlibabaAlihealthMedicalItemPublishAPIRequest) GetRequest1() *ItemPublishRequest {
    return r._request1
}
