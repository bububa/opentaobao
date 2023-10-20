package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkordersyncwithitemAPIRequest 订单和商品同步接口 API请求
// alibaba.wdk.order.sync.withitem
//
// 轻pos,将订单和商品的信息一起传到盒马这边，进行创单和添加商品处理。
type AlibabawdkordersyncwithitemAPIRequest struct {
	model.Params
	// 商家传过来的交易和商品信息
	_posOrderAndItemSync *PosOrderAndItemSyncDo
}

// NewAlibabawdkordersyncwithitemRequest 初始化AlibabawdkordersyncwithitemAPIRequest对象
func NewAlibabawdkordersyncwithitemRequest() *AlibabawdkordersyncwithitemAPIRequest {
	return &AlibabawdkordersyncwithitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkordersyncwithitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.sync.withitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkordersyncwithitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkordersyncwithitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosOrderAndItemSync is PosOrderAndItemSync Setter
// 商家传过来的交易和商品信息
func (r *AlibabawdkordersyncwithitemAPIRequest) SetPosOrderAndItemSync(_posOrderAndItemSync *PosOrderAndItemSyncDo) error {
	r._posOrderAndItemSync = _posOrderAndItemSync
	r.Set("pos_order_and_item_sync", _posOrderAndItemSync)
	return nil
}

// GetPosOrderAndItemSync PosOrderAndItemSync Getter
func (r AlibabawdkordersyncwithitemAPIRequest) GetPosOrderAndItemSync() *PosOrderAndItemSyncDo {
	return r._posOrderAndItemSync
}
