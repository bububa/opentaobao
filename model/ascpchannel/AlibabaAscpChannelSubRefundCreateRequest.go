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
type AlibabaAscpChannelSubRefundCreateRequest struct {
    model.Params
    // 子单退款创建请求
    _subRefundCreateReq   *ExternalCreateRefundOrderDetailRequest
}

// 初始化AlibabaAscpChannelSubRefundCreateRequest对象
func NewAlibabaAscpChannelSubRefundCreateRequest() *AlibabaAscpChannelSubRefundCreateRequest{
    return &AlibabaAscpChannelSubRefundCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSubRefundCreateRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.sub.refund.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSubRefundCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubRefundCreateReq Setter
// 子单退款创建请求
func (r *AlibabaAscpChannelSubRefundCreateRequest) SetSubRefundCreateReq(_subRefundCreateReq *ExternalCreateRefundOrderDetailRequest) error {
    r._subRefundCreateReq = _subRefundCreateReq
    r.Set("sub_refund_create_req", _subRefundCreateReq)
    return nil
}

// SubRefundCreateReq Getter
func (r AlibabaAscpChannelSubRefundCreateRequest) GetSubRefundCreateReq() *ExternalCreateRefundOrderDetailRequest {
    return r._subRefundCreateReq
}
