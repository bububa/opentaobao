package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达外部商户老pos机产生的订单同步进淘鲜达 API请求
alibaba.wdk.oldpos.order.create

淘鲜达外部商户老pos机产生的订单同步进淘鲜达
*/
type AlibabaWdkOldposOrderCreateRequest struct {
    model.Params
    // 入参
    posOrderCreateRequest   *PosOrderCreateRequest
}

// 初始化AlibabaWdkOldposOrderCreateRequest对象
func NewAlibabaWdkOldposOrderCreateRequest() *AlibabaWdkOldposOrderCreateRequest{
    return &AlibabaWdkOldposOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOldposOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.oldpos.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOldposOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosOrderCreateRequest Setter
// 入参
func (r *AlibabaWdkOldposOrderCreateRequest) SetPosOrderCreateRequest(posOrderCreateRequest *PosOrderCreateRequest) error {
    r.posOrderCreateRequest = posOrderCreateRequest
    r.Set("pos_order_create_request", posOrderCreateRequest)
    return nil
}

// PosOrderCreateRequest Getter
func (r AlibabaWdkOldposOrderCreateRequest) GetPosOrderCreateRequest() *PosOrderCreateRequest {
    return r.posOrderCreateRequest
}
