package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ASCP渠道中心销售单创建接口 APIRequest
tmall.ascp.orders.sale.create

ASCP渠道中心销售单创建接口
*/
type TmallAscpOrdersSaleCreateRequest struct {
    model.Params

    // 请求对象
    channelOrderRequest   *CreateChannelOrderRequest 

}

func NewTmallAscpOrdersSaleCreateRequest() *TmallAscpOrdersSaleCreateRequest{
    return &TmallAscpOrdersSaleCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallAscpOrdersSaleCreateRequest) GetApiMethodName() string {
    return "tmall.ascp.orders.sale.create"
}

func (r TmallAscpOrdersSaleCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallAscpOrdersSaleCreateRequest) SetChannelOrderRequest(channelOrderRequest *CreateChannelOrderRequest) error {
    r.channelOrderRequest = channelOrderRequest
    r.Set("channel_order_request", channelOrderRequest)
    return nil
}

func (r TmallAscpOrdersSaleCreateRequest) GetChannelOrderRequest() *CreateChannelOrderRequest {
    return r.channelOrderRequest
}

