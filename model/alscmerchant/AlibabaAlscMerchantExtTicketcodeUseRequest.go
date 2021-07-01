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
type AlibabaAlscMerchantExtTicketcodeUseAPIRequest struct {
    model.Params
    // 外部券使用请求
    _useRequest   *ExternalTicketUseRequest
}

// 初始化AlibabaAlscMerchantExtTicketcodeUseAPIRequest对象
func NewAlibabaAlscMerchantExtTicketcodeUseRequest() *AlibabaAlscMerchantExtTicketcodeUseAPIRequest{
    return &AlibabaAlscMerchantExtTicketcodeUseAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscMerchantExtTicketcodeUseAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.merchant.ext.ticketcode.use"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscMerchantExtTicketcodeUseAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UseRequest Setter
// 外部券使用请求
func (r *AlibabaAlscMerchantExtTicketcodeUseAPIRequest) SetUseRequest(_useRequest *ExternalTicketUseRequest) error {
    r._useRequest = _useRequest
    r.Set("use_request", _useRequest)
    return nil
}

// UseRequest Getter
func (r AlibabaAlscMerchantExtTicketcodeUseAPIRequest) GetUseRequest() *ExternalTicketUseRequest {
    return r._useRequest
}
