package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通发货单取消 API请求
alibaba.lst.shiporder.cancel

通过该接口可以取消零售通运保保发货单，并处理相关业务流程。
*/
type AlibabaLstShiporderCancelRequest struct {
    model.Params
    // 取消原因
    _reason   string
    // 订单号
    _outOrderId   string
    // 需要退款的明细ID
    _detailOrderIds   []string
}

// 初始化AlibabaLstShiporderCancelRequest对象
func NewAlibabaLstShiporderCancelRequest() *AlibabaLstShiporderCancelRequest{
    return &AlibabaLstShiporderCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstShiporderCancelRequest) GetApiMethodName() string {
    return "alibaba.lst.shiporder.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstShiporderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Reason Setter
// 取消原因
func (r *AlibabaLstShiporderCancelRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r AlibabaLstShiporderCancelRequest) GetReason() string {
    return r._reason
}
// OutOrderId Setter
// 订单号
func (r *AlibabaLstShiporderCancelRequest) SetOutOrderId(_outOrderId string) error {
    r._outOrderId = _outOrderId
    r.Set("out_order_id", _outOrderId)
    return nil
}

// OutOrderId Getter
func (r AlibabaLstShiporderCancelRequest) GetOutOrderId() string {
    return r._outOrderId
}
// DetailOrderIds Setter
// 需要退款的明细ID
func (r *AlibabaLstShiporderCancelRequest) SetDetailOrderIds(_detailOrderIds []string) error {
    r._detailOrderIds = _detailOrderIds
    r.Set("detail_order_ids", _detailOrderIds)
    return nil
}

// DetailOrderIds Getter
func (r AlibabaLstShiporderCancelRequest) GetDetailOrderIds() []string {
    return r._detailOrderIds
}
