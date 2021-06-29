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
type AlibabaWdkOldposRefundCreateRequest struct {
    model.Params
    // 入参
    posRefundCreateRequest   *PosRefundCreateRequest
}

// 初始化AlibabaWdkOldposRefundCreateRequest对象
func NewAlibabaWdkOldposRefundCreateRequest() *AlibabaWdkOldposRefundCreateRequest{
    return &AlibabaWdkOldposRefundCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOldposRefundCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.oldpos.refund.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOldposRefundCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosRefundCreateRequest Setter
// 入参
func (r *AlibabaWdkOldposRefundCreateRequest) SetPosRefundCreateRequest(posRefundCreateRequest *PosRefundCreateRequest) error {
    r.posRefundCreateRequest = posRefundCreateRequest
    r.Set("pos_refund_create_request", posRefundCreateRequest)
    return nil
}

// PosRefundCreateRequest Getter
func (r AlibabaWdkOldposRefundCreateRequest) GetPosRefundCreateRequest() *PosRefundCreateRequest {
    return r.posRefundCreateRequest
}
