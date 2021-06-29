package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】退票申请单详情 API请求
taobao.alitrip.seller.refund.get

查询退票申请单详情
*/
type TaobaoAlitripSellerRefundGetRequest struct {
    model.Params
    // 申请单ID
    _applyId   int64
}

// 初始化TaobaoAlitripSellerRefundGetRequest对象
func NewTaobaoAlitripSellerRefundGetRequest() *TaobaoAlitripSellerRefundGetRequest{
    return &TaobaoAlitripSellerRefundGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refund.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请单ID
func (r *TaobaoAlitripSellerRefundGetRequest) SetApplyId(_applyId int64) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripSellerRefundGetRequest) GetApplyId() int64 {
    return r._applyId
}
