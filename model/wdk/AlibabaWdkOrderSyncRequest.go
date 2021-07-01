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
type AlibabaWdkOrderSyncAPIRequest struct {
    model.Params
    // 订单
    _receiptOrder   *ReceiptOrderDO
}

// 初始化AlibabaWdkOrderSyncAPIRequest对象
func NewAlibabaWdkOrderSyncRequest() *AlibabaWdkOrderSyncAPIRequest{
    return &AlibabaWdkOrderSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReceiptOrder Setter
// 订单
func (r *AlibabaWdkOrderSyncAPIRequest) SetReceiptOrder(_receiptOrder *ReceiptOrderDO) error {
    r._receiptOrder = _receiptOrder
    r.Set("receipt_order", _receiptOrder)
    return nil
}

// ReceiptOrder Getter
func (r AlibabaWdkOrderSyncAPIRequest) GetReceiptOrder() *ReceiptOrderDO {
    return r._receiptOrder
}
