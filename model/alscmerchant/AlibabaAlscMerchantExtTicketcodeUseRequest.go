package alscmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部核销服务 API请求
alibaba.alsc.merchant.ext.ticketcode.use

外部核销服务
*/
type AlibabaAlscMerchantExtTicketcodeUseRequest struct {
    model.Params
    // 外部券使用请求
    useRequest   *ExternalTicketUseRequest
}

// 初始化AlibabaAlscMerchantExtTicketcodeUseRequest对象
func NewAlibabaAlscMerchantExtTicketcodeUseRequest() *AlibabaAlscMerchantExtTicketcodeUseRequest{
    return &AlibabaAlscMerchantExtTicketcodeUseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscMerchantExtTicketcodeUseRequest) GetApiMethodName() string {
    return "alibaba.alsc.merchant.ext.ticketcode.use"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscMerchantExtTicketcodeUseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UseRequest Setter
// 外部券使用请求
func (r *AlibabaAlscMerchantExtTicketcodeUseRequest) SetUseRequest(useRequest *ExternalTicketUseRequest) error {
    r.useRequest = useRequest
    r.Set("use_request", useRequest)
    return nil
}

// UseRequest Getter
func (r AlibabaAlscMerchantExtTicketcodeUseRequest) GetUseRequest() *ExternalTicketUseRequest {
    return r.useRequest
}
