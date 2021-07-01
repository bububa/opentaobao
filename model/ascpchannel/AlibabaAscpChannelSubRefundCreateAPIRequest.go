package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘外分销逆向创单（子单退） API请求
alibaba.ascp.channel.sub.refund.create

淘外分销逆向创单（子单退）
*/
type AlibabaAscpChannelSubRefundCreateAPIRequest struct {
    model.Params
    // 子单退款创建请求
    _subRefundCreateReq   *ExternalCreateRefundOrderDetailRequest
}

// 初始化AlibabaAscpChannelSubRefundCreateAPIRequest对象
func NewAlibabaAscpChannelSubRefundCreateRequest() *AlibabaAscpChannelSubRefundCreateAPIRequest{
    return &AlibabaAscpChannelSubRefundCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSubRefundCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.sub.refund.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSubRefundCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubRefundCreateReq Setter
// 子单退款创建请求
func (r *AlibabaAscpChannelSubRefundCreateAPIRequest) SetSubRefundCreateReq(_subRefundCreateReq *ExternalCreateRefundOrderDetailRequest) error {
    r._subRefundCreateReq = _subRefundCreateReq
    r.Set("sub_refund_create_req", _subRefundCreateReq)
    return nil
}

// SubRefundCreateReq Getter
func (r AlibabaAscpChannelSubRefundCreateAPIRequest) GetSubRefundCreateReq() *ExternalCreateRefundOrderDetailRequest {
    return r._subRefundCreateReq
}
