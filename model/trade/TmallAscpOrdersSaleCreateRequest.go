package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ASCP渠道中心销售单创建接口 API请求
tmall.ascp.orders.sale.create

ASCP渠道中心销售单创建接口
*/
type TmallAscpOrdersSaleCreateRequest struct {
    model.Params
    // 请求对象
    _channelOrderRequest   *CreateChannelOrderRequest
}

// 初始化TmallAscpOrdersSaleCreateRequest对象
func NewTmallAscpOrdersSaleCreateRequest() *TmallAscpOrdersSaleCreateRequest{
    return &TmallAscpOrdersSaleCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAscpOrdersSaleCreateRequest) GetApiMethodName() string {
    return "tmall.ascp.orders.sale.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallAscpOrdersSaleCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelOrderRequest Setter
// 请求对象
func (r *TmallAscpOrdersSaleCreateRequest) SetChannelOrderRequest(_channelOrderRequest *CreateChannelOrderRequest) error {
    r._channelOrderRequest = _channelOrderRequest
    r.Set("channel_order_request", _channelOrderRequest)
    return nil
}

// ChannelOrderRequest Getter
func (r TmallAscpOrdersSaleCreateRequest) GetChannelOrderRequest() *CreateChannelOrderRequest {
    return r._channelOrderRequest
}
