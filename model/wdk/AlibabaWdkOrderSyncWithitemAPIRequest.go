package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderSyncWithitemAPIRequest 订单和商品同步接口 API请求
// alibaba.wdk.order.sync.withitem
//
// 轻pos,将订单和商品的信息一起传到盒马这边，进行创单和添加商品处理。
type AlibabaWdkOrderSyncWithitemAPIRequest struct {
	model.Params
	// 商家传过来的交易和商品信息
	_posOrderAndItemSync *PosOrderAndItemSyncDo
}

// NewAlibabaWdkOrderSyncWithitemRequest 初始化AlibabaWdkOrderSyncWithitemAPIRequest对象
func NewAlibabaWdkOrderSyncWithitemRequest() *AlibabaWdkOrderSyncWithitemAPIRequest {
	return &AlibabaWdkOrderSyncWithitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderSyncWithitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.sync.withitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderSyncWithitemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPosOrderAndItemSync is PosOrderAndItemSync Setter
// 商家传过来的交易和商品信息
func (r *AlibabaWdkOrderSyncWithitemAPIRequest) SetPosOrderAndItemSync(_posOrderAndItemSync *PosOrderAndItemSyncDo) error {
	r._posOrderAndItemSync = _posOrderAndItemSync
	r.Set("pos_order_and_item_sync", _posOrderAndItemSync)
	return nil
}

// GetPosOrderAndItemSync PosOrderAndItemSync Getter
func (r AlibabaWdkOrderSyncWithitemAPIRequest) GetPosOrderAndItemSync() *PosOrderAndItemSyncDo {
	return r._posOrderAndItemSync
}
