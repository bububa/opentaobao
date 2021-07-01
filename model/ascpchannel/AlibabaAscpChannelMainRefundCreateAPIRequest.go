package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘外分销逆向创单（未发货整单退） API请求
alibaba.ascp.channel.main.refund.create

淘外分销解决方案--订单--逆向创单（未发货整单退）
*/
type AlibabaAscpChannelMainRefundCreateAPIRequest struct {
    model.Params
    // 逆向单创建请求
    _refundCreateRequest   *ExternalCreateRefundOrderRequest
}

// 初始化AlibabaAscpChannelMainRefundCreateAPIRequest对象
func NewAlibabaAscpChannelMainRefundCreateRequest() *AlibabaAscpChannelMainRefundCreateAPIRequest{
    return &AlibabaAscpChannelMainRefundCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelMainRefundCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.main.refund.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelMainRefundCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundCreateRequest Setter
// 逆向单创建请求
func (r *AlibabaAscpChannelMainRefundCreateAPIRequest) SetRefundCreateRequest(_refundCreateRequest *ExternalCreateRefundOrderRequest) error {
    r._refundCreateRequest = _refundCreateRequest
    r.Set("refund_create_request", _refundCreateRequest)
    return nil
}

// RefundCreateRequest Getter
func (r AlibabaAscpChannelMainRefundCreateAPIRequest) GetRefundCreateRequest() *ExternalCreateRefundOrderRequest {
    return r._refundCreateRequest
}
