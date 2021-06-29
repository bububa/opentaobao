package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘外分销逆向创单（子单退） APIRequest
alibaba.ascp.channel.sub.refund.create

淘外分销逆向创单（子单退）
*/
type AlibabaAscpChannelSubRefundCreateRequest struct {
    model.Params

    // 子单退款创建请求
    subRefundCreateReq   *ExternalCreateRefundOrderDetailRequest 

}

func NewAlibabaAscpChannelSubRefundCreateRequest() *AlibabaAscpChannelSubRefundCreateRequest{
    return &AlibabaAscpChannelSubRefundCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpChannelSubRefundCreateRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.sub.refund.create"
}

func (r AlibabaAscpChannelSubRefundCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpChannelSubRefundCreateRequest) SetSubRefundCreateReq(subRefundCreateReq *ExternalCreateRefundOrderDetailRequest) error {
    r.subRefundCreateReq = subRefundCreateReq
    r.Set("sub_refund_create_req", subRefundCreateReq)
    return nil
}

func (r AlibabaAscpChannelSubRefundCreateRequest) GetSubRefundCreateReq() *ExternalCreateRefundOrderDetailRequest {
    return r.subRefundCreateReq
}

