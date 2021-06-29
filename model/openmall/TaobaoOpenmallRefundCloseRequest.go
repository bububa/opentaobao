package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关闭OpenMall退款单 API请求
taobao.openmall.refund.close

关闭OpenMall退款单
*/
type TaobaoOpenmallRefundCloseRequest struct {
    model.Params
    // 渠道
    _distributor   string
    // 退款ID
    _refundId   int64
}

// 初始化TaobaoOpenmallRefundCloseRequest对象
func NewTaobaoOpenmallRefundCloseRequest() *TaobaoOpenmallRefundCloseRequest{
    return &TaobaoOpenmallRefundCloseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundCloseRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.close"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 渠道
func (r *TaobaoOpenmallRefundCloseRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundCloseRequest) GetDistributor() string {
    return r._distributor
}
// RefundId Setter
// 退款ID
func (r *TaobaoOpenmallRefundCloseRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoOpenmallRefundCloseRequest) GetRefundId() int64 {
    return r._refundId
}
