package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单条退款跟踪详情 API请求
taobao.jds.refund.traces.get

获取聚石塔数据共享的交易全链路的退款信息
*/
type TaobaoJdsRefundTracesGetRequest struct {
    model.Params
    // 淘宝的退款编号
    refundId   int64
}

// 初始化TaobaoJdsRefundTracesGetRequest对象
func NewTaobaoJdsRefundTracesGetRequest() *TaobaoJdsRefundTracesGetRequest{
    return &TaobaoJdsRefundTracesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJdsRefundTracesGetRequest) GetApiMethodName() string {
    return "taobao.jds.refund.traces.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJdsRefundTracesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 淘宝的退款编号
func (r *TaobaoJdsRefundTracesGetRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoJdsRefundTracesGetRequest) GetRefundId() int64 {
    return r.refundId
}
