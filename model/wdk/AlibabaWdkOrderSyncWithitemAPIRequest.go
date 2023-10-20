package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkOrderSyncWithitemAPIRequest) Reset() {
	r._posOrderAndItemSync = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderSyncWithitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.sync.withitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderSyncWithitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkOrderSyncWithitemAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkOrderSyncWithitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkOrderSyncWithitemRequest()
	},
}

// GetAlibabaWdkOrderSyncWithitemRequest 从 sync.Pool 获取 AlibabaWdkOrderSyncWithitemAPIRequest
func GetAlibabaWdkOrderSyncWithitemAPIRequest() *AlibabaWdkOrderSyncWithitemAPIRequest {
	return poolAlibabaWdkOrderSyncWithitemAPIRequest.Get().(*AlibabaWdkOrderSyncWithitemAPIRequest)
}

// ReleaseAlibabaWdkOrderSyncWithitemAPIRequest 将 AlibabaWdkOrderSyncWithitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkOrderSyncWithitemAPIRequest(v *AlibabaWdkOrderSyncWithitemAPIRequest) {
	v.Reset()
	poolAlibabaWdkOrderSyncWithitemAPIRequest.Put(v)
}
