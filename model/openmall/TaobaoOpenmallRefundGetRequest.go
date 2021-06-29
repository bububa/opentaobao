package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取OpenMall退款单详情 API请求
taobao.openmall.refund.get

获取OpenMall退款单详情
*/
type TaobaoOpenmallRefundGetRequest struct {
    model.Params
    // 渠道商身份
    _distributor   string
    // 退款单ID
    _refundId   int64
}

// 初始化TaobaoOpenmallRefundGetRequest对象
func NewTaobaoOpenmallRefundGetRequest() *TaobaoOpenmallRefundGetRequest{
    return &TaobaoOpenmallRefundGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundGetRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 渠道商身份
func (r *TaobaoOpenmallRefundGetRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundGetRequest) GetDistributor() string {
    return r._distributor
}
// RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundGetRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoOpenmallRefundGetRequest) GetRefundId() int64 {
    return r._refundId
}
