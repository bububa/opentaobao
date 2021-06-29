package ascpqcc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
魅力惠样品解除父子商品关系 API请求
alibaba.ascp.qcc.sample.cancel.item.relation

品控中心魅力惠样品解除父子商品关系
*/
type AlibabaAscpQccSampleCancelItemRelationRequest struct {
    model.Params
    // 请求参数对象
    cancelRequest   *CancelSampleRelationRequest
}

// 初始化AlibabaAscpQccSampleCancelItemRelationRequest对象
func NewAlibabaAscpQccSampleCancelItemRelationRequest() *AlibabaAscpQccSampleCancelItemRelationRequest{
    return &AlibabaAscpQccSampleCancelItemRelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpQccSampleCancelItemRelationRequest) GetApiMethodName() string {
    return "alibaba.ascp.qcc.sample.cancel.item.relation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpQccSampleCancelItemRelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CancelRequest Setter
// 请求参数对象
func (r *AlibabaAscpQccSampleCancelItemRelationRequest) SetCancelRequest(cancelRequest *CancelSampleRelationRequest) error {
    r.cancelRequest = cancelRequest
    r.Set("cancel_request", cancelRequest)
    return nil
}

// CancelRequest Getter
func (r AlibabaAscpQccSampleCancelItemRelationRequest) GetCancelRequest() *CancelSampleRelationRequest {
    return r.cancelRequest
}
