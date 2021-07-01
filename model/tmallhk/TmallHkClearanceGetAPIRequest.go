package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫国际-清关材料查询 API请求
tmall.hk.clearance.get

提供订单收货人身份信息查询功能。
*/
type TmallHkClearanceGetAPIRequest struct {
    model.Params
    // 天猫国际订单号
    _orderId   int64
    // 是否需要身份证图片，不需要可以缩短接口响应时间
    _needImage   bool
}

// 初始化TmallHkClearanceGetAPIRequest对象
func NewTmallHkClearanceGetRequest() *TmallHkClearanceGetAPIRequest{
    return &TmallHkClearanceGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallHkClearanceGetAPIRequest) GetApiMethodName() string {
    return "tmall.hk.clearance.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallHkClearanceGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 天猫国际订单号
func (r *TmallHkClearanceGetAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallHkClearanceGetAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// NeedImage Setter
// 是否需要身份证图片，不需要可以缩短接口响应时间
func (r *TmallHkClearanceGetAPIRequest) SetNeedImage(_needImage bool) error {
    r._needImage = _needImage
    r.Set("need_image", _needImage)
    return nil
}

// NeedImage Getter
func (r TmallHkClearanceGetAPIRequest) GetNeedImage() bool {
    return r._needImage
}
