package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口外部商户老pos机产生的退款单同步进盒马 API请求
alibaba.wdk.oldpos.refund.create

淘鲜达外部商户老pos机产生的退款单同步进淘鲜达
*/
type AlibabaWdkOldposRefundCreateAPIRequest struct {
    model.Params
    // 入参
    _posRefundCreateRequest   *PosRefundCreateRequest
}

// 初始化AlibabaWdkOldposRefundCreateAPIRequest对象
func NewAlibabaWdkOldposRefundCreateRequest() *AlibabaWdkOldposRefundCreateAPIRequest{
    return &AlibabaWdkOldposRefundCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOldposRefundCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.oldpos.refund.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOldposRefundCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosRefundCreateRequest Setter
// 入参
func (r *AlibabaWdkOldposRefundCreateAPIRequest) SetPosRefundCreateRequest(_posRefundCreateRequest *PosRefundCreateRequest) error {
    r._posRefundCreateRequest = _posRefundCreateRequest
    r.Set("pos_refund_create_request", _posRefundCreateRequest)
    return nil
}

// PosRefundCreateRequest Getter
func (r AlibabaWdkOldposRefundCreateAPIRequest) GetPosRefundCreateRequest() *PosRefundCreateRequest {
    return r._posRefundCreateRequest
}
