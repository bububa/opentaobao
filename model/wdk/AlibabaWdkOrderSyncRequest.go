package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口外部订单同步 API请求
alibaba.wdk.order.sync

外部商户使用自助POS下单订单同步到五道口
*/
type AlibabaWdkOrderSyncRequest struct {
    model.Params
    // 订单
    receiptOrder   *ReceiptOrderDO
}

// 初始化AlibabaWdkOrderSyncRequest对象
func NewAlibabaWdkOrderSyncRequest() *AlibabaWdkOrderSyncRequest{
    return &AlibabaWdkOrderSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderSyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReceiptOrder Setter
// 订单
func (r *AlibabaWdkOrderSyncRequest) SetReceiptOrder(receiptOrder *ReceiptOrderDO) error {
    r.receiptOrder = receiptOrder
    r.Set("receipt_order", receiptOrder)
    return nil
}

// ReceiptOrder Getter
func (r AlibabaWdkOrderSyncRequest) GetReceiptOrder() *ReceiptOrderDO {
    return r.receiptOrder
}
