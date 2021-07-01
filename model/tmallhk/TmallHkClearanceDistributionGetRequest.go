package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销供应商获取清关材料 API请求
tmall.hk.clearance.distribution.get

供销体系下，提供供应商可以直接获取其订单身份证信息的接口，以使其完成清关。
*/
type TmallHkClearanceDistributionGetAPIRequest struct {
    model.Params
    // 订单号
    _orderId   int64
    // 是否需要身份证图片，不需要可以缩短接口响应时间
    _needImage   bool
}

// 初始化TmallHkClearanceDistributionGetAPIRequest对象
func NewTmallHkClearanceDistributionGetRequest() *TmallHkClearanceDistributionGetAPIRequest{
    return &TmallHkClearanceDistributionGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallHkClearanceDistributionGetAPIRequest) GetApiMethodName() string {
    return "tmall.hk.clearance.distribution.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallHkClearanceDistributionGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单号
func (r *TmallHkClearanceDistributionGetAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallHkClearanceDistributionGetAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// NeedImage Setter
// 是否需要身份证图片，不需要可以缩短接口响应时间
func (r *TmallHkClearanceDistributionGetAPIRequest) SetNeedImage(_needImage bool) error {
    r._needImage = _needImage
    r.Set("need_image", _needImage)
    return nil
}

// NeedImage Getter
func (r TmallHkClearanceDistributionGetAPIRequest) GetNeedImage() bool {
    return r._needImage
}
