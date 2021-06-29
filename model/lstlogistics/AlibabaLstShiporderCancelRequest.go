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
    reason   string
    // 订单号
    outOrderId   string
    // 需要退款的明细ID
    detailOrderIds   []string
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
func (r *AlibabaLstShiporderCancelRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

// Reason Getter
func (r AlibabaLstShiporderCancelRequest) GetReason() string {
    return r.reason
}
// OutOrderId Setter
// 订单号
func (r *AlibabaLstShiporderCancelRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r AlibabaLstShiporderCancelRequest) GetOutOrderId() string {
    return r.outOrderId
}
// DetailOrderIds Setter
// 需要退款的明细ID
func (r *AlibabaLstShiporderCancelRequest) SetDetailOrderIds(detailOrderIds []string) error {
    r.detailOrderIds = detailOrderIds
    r.Set("detail_order_ids", detailOrderIds)
    return nil
}

// DetailOrderIds Getter
func (r AlibabaLstShiporderCancelRequest) GetDetailOrderIds() []string {
    return r.detailOrderIds
}
