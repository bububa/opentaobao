package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品上下架 APIRequest
alibaba.alihealth.medical.item.status

医生通三方机构平台进行服务商品上下架操作
*/
type AlibabaAlihealthMedicalItemStatusRequest struct {
    model.Params

    // 请求入参
    shelfrequest   *ThirdAgencyUpDownShelfRequest 

}

func NewAlibabaAlihealthMedicalItemStatusRequest() *AlibabaAlihealthMedicalItemStatusRequest{
    return &AlibabaAlihealthMedicalItemStatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalItemStatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.item.status"
}

func (r AlibabaAlihealthMedicalItemStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalItemStatusRequest) SetShelfrequest(shelfrequest *ThirdAgencyUpDownShelfRequest) error {
    r.shelfrequest = shelfrequest
    r.Set("shelfrequest", shelfrequest)
    return nil
}

func (r AlibabaAlihealthMedicalItemStatusRequest) GetShelfrequest() *ThirdAgencyUpDownShelfRequest {
    return r.shelfrequest
}

