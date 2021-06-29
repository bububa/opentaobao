package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品上下架 API请求
alibaba.alihealth.medical.item.status

医生通三方机构平台进行服务商品上下架操作
*/
type AlibabaAlihealthMedicalItemStatusRequest struct {
    model.Params
    // 请求入参
    _shelfrequest   *ThirdAgencyUpDownShelfRequest
}

// 初始化AlibabaAlihealthMedicalItemStatusRequest对象
func NewAlibabaAlihealthMedicalItemStatusRequest() *AlibabaAlihealthMedicalItemStatusRequest{
    return &AlibabaAlihealthMedicalItemStatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalItemStatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.item.status"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalItemStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Shelfrequest Setter
// 请求入参
func (r *AlibabaAlihealthMedicalItemStatusRequest) SetShelfrequest(_shelfrequest *ThirdAgencyUpDownShelfRequest) error {
    r._shelfrequest = _shelfrequest
    r.Set("shelfrequest", _shelfrequest)
    return nil
}

// Shelfrequest Getter
func (r AlibabaAlihealthMedicalItemStatusRequest) GetShelfrequest() *ThirdAgencyUpDownShelfRequest {
    return r._shelfrequest
}
