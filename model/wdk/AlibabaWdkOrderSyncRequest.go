package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口外部订单同步 APIRequest
alibaba.wdk.order.sync

外部商户使用自助POS下单订单同步到五道口
*/
type AlibabaWdkOrderSyncRequest struct {
    model.Params

    // 订单
    receiptOrder   *ReceiptOrderDO 

}

func NewAlibabaWdkOrderSyncRequest() *AlibabaWdkOrderSyncRequest{
    return &AlibabaWdkOrderSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkOrderSyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.sync"
}

func (r AlibabaWdkOrderSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkOrderSyncRequest) SetReceiptOrder(receiptOrder *ReceiptOrderDO) error {
    r.receiptOrder = receiptOrder
    r.Set("receipt_order", receiptOrder)
    return nil
}

func (r AlibabaWdkOrderSyncRequest) GetReceiptOrder() *ReceiptOrderDO {
    return r.receiptOrder
}

