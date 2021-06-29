package alscmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部核销服务 APIRequest
alibaba.alsc.merchant.ext.ticketcode.use

外部核销服务
*/
type AlibabaAlscMerchantExtTicketcodeUseRequest struct {
    model.Params

    // 外部券使用请求
    useRequest   *ExternalTicketUseRequest 

}

func NewAlibabaAlscMerchantExtTicketcodeUseRequest() *AlibabaAlscMerchantExtTicketcodeUseRequest{
    return &AlibabaAlscMerchantExtTicketcodeUseRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscMerchantExtTicketcodeUseRequest) GetApiMethodName() string {
    return "alibaba.alsc.merchant.ext.ticketcode.use"
}

func (r AlibabaAlscMerchantExtTicketcodeUseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscMerchantExtTicketcodeUseRequest) SetUseRequest(useRequest *ExternalTicketUseRequest) error {
    r.useRequest = useRequest
    r.Set("use_request", useRequest)
    return nil
}

func (r AlibabaAlscMerchantExtTicketcodeUseRequest) GetUseRequest() *ExternalTicketUseRequest {
    return r.useRequest
}

