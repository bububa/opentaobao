package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单和商品同步接口 API请求
alibaba.wdk.order.sync.withitem

轻pos,将订单和商品的信息一起传到盒马这边，进行创单和添加商品处理。
*/
type AlibabaWdkOrderSyncWithitemAPIRequest struct {
    model.Params
    // 商家传过来的交易和商品信息
    _posOrderAndItemSync   *PosOrderAndItemSyncDO
}

// 初始化AlibabaWdkOrderSyncWithitemAPIRequest对象
func NewAlibabaWdkOrderSyncWithitemRequest() *AlibabaWdkOrderSyncWithitemAPIRequest{
    return &AlibabaWdkOrderSyncWithitemAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderSyncWithitemAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.sync.withitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderSyncWithitemAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosOrderAndItemSync Setter
// 商家传过来的交易和商品信息
func (r *AlibabaWdkOrderSyncWithitemAPIRequest) SetPosOrderAndItemSync(_posOrderAndItemSync *PosOrderAndItemSyncDO) error {
    r._posOrderAndItemSync = _posOrderAndItemSync
    r.Set("pos_order_and_item_sync", _posOrderAndItemSync)
    return nil
}

// PosOrderAndItemSync Getter
func (r AlibabaWdkOrderSyncWithitemAPIRequest) GetPosOrderAndItemSync() *PosOrderAndItemSyncDO {
    return r._posOrderAndItemSync
}
