package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销供应商获取清关材料 APIRequest
tmall.hk.clearance.distribution.get

供销体系下，提供供应商可以直接获取其订单身份证信息的接口，以使其完成清关。
*/
type TmallHkClearanceDistributionGetRequest struct {
    model.Params

    // 订单号
    orderId   int64 

    // 是否需要身份证图片，不需要可以缩短接口响应时间
    needImage   bool 

}

func NewTmallHkClearanceDistributionGetRequest() *TmallHkClearanceDistributionGetRequest{
    return &TmallHkClearanceDistributionGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallHkClearanceDistributionGetRequest) GetApiMethodName() string {
    return "tmall.hk.clearance.distribution.get"
}

func (r TmallHkClearanceDistributionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallHkClearanceDistributionGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TmallHkClearanceDistributionGetRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TmallHkClearanceDistributionGetRequest) SetNeedImage(needImage bool) error {
    r.needImage = needImage
    r.Set("need_image", needImage)
    return nil
}

func (r TmallHkClearanceDistributionGetRequest) GetNeedImage() bool {
    return r.needImage
}

