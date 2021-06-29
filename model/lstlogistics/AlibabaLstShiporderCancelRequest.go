package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通发货单取消 APIRequest
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

func NewAlibabaLstShiporderCancelRequest() *AlibabaLstShiporderCancelRequest{
    return &AlibabaLstShiporderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstShiporderCancelRequest) GetApiMethodName() string {
    return "alibaba.lst.shiporder.cancel"
}

func (r AlibabaLstShiporderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstShiporderCancelRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r AlibabaLstShiporderCancelRequest) GetReason() string {
    return r.reason
}

func (r *AlibabaLstShiporderCancelRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

func (r AlibabaLstShiporderCancelRequest) GetOutOrderId() string {
    return r.outOrderId
}

func (r *AlibabaLstShiporderCancelRequest) SetDetailOrderIds(detailOrderIds []string) error {
    r.detailOrderIds = detailOrderIds
    r.Set("detail_order_ids", detailOrderIds)
    return nil
}

func (r AlibabaLstShiporderCancelRequest) GetDetailOrderIds() []string {
    return r.detailOrderIds
}

