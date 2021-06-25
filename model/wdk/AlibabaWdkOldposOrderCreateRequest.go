package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达外部商户老pos机产生的订单同步进淘鲜达 APIRequest
alibaba.wdk.oldpos.order.create

淘鲜达外部商户老pos机产生的订单同步进淘鲜达
*/
type AlibabaWdkOldposOrderCreateRequest struct {
    model.Params

    // 入参
    posOrderCreateRequest   *PosOrderCreateRequest 

}

func NewAlibabaWdkOldposOrderCreateRequest() *AlibabaWdkOldposOrderCreateRequest{
    return &AlibabaWdkOldposOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkOldposOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.oldpos.order.create"
}

func (r AlibabaWdkOldposOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkOldposOrderCreateRequest) SetPosOrderCreateRequest(posOrderCreateRequest *PosOrderCreateRequest) error {
    r.posOrderCreateRequest = posOrderCreateRequest
    r.Set("pos_order_create_request", posOrderCreateRequest)
    return nil
}

func (r AlibabaWdkOldposOrderCreateRequest) GetPosOrderCreateRequest() *PosOrderCreateRequest {
    return r.posOrderCreateRequest
}

