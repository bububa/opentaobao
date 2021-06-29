package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘外分销逆向创单（未发货整单退） APIRequest
alibaba.ascp.channel.main.refund.create

淘外分销解决方案--订单--逆向创单（未发货整单退）
*/
type AlibabaAscpChannelMainRefundCreateRequest struct {
    model.Params

    // 逆向单创建请求
    refundCreateRequest   *ExternalCreateRefundOrderRequest 

}

func NewAlibabaAscpChannelMainRefundCreateRequest() *AlibabaAscpChannelMainRefundCreateRequest{
    return &AlibabaAscpChannelMainRefundCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpChannelMainRefundCreateRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.main.refund.create"
}

func (r AlibabaAscpChannelMainRefundCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpChannelMainRefundCreateRequest) SetRefundCreateRequest(refundCreateRequest *ExternalCreateRefundOrderRequest) error {
    r.refundCreateRequest = refundCreateRequest
    r.Set("refund_create_request", refundCreateRequest)
    return nil
}

func (r AlibabaAscpChannelMainRefundCreateRequest) GetRefundCreateRequest() *ExternalCreateRefundOrderRequest {
    return r.refundCreateRequest
}

