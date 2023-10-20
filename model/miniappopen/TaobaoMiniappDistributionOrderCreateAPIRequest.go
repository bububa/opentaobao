package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionOrderCreateAPIRequest 创建小程序投放计划 API请求
// taobao.miniapp.distribution.order.create
//
// 帮助商家，创建小程序的投放计划。该api，仅针对特定场景开放，目前仅支持客服场景，具体支持的场景列表请联系技术支持或业务对接人确认。
type TaobaoMiniappDistributionOrderCreateAPIRequest struct {
	model.Params
	// 投放计划信息
	_orderRequest *DistributionOrderSaveOpenRequest
}

// NewTaobaoMiniappDistributionOrderCreateRequest 初始化TaobaoMiniappDistributionOrderCreateAPIRequest对象
func NewTaobaoMiniappDistributionOrderCreateRequest() *TaobaoMiniappDistributionOrderCreateAPIRequest {
	return &TaobaoMiniappDistributionOrderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappDistributionOrderCreateAPIRequest) Reset() {
	r._orderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionOrderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappDistributionOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRequest is OrderRequest Setter
// 投放计划信息
func (r *TaobaoMiniappDistributionOrderCreateAPIRequest) SetOrderRequest(_orderRequest *DistributionOrderSaveOpenRequest) error {
	r._orderRequest = _orderRequest
	r.Set("order_request", _orderRequest)
	return nil
}

// GetOrderRequest OrderRequest Getter
func (r TaobaoMiniappDistributionOrderCreateAPIRequest) GetOrderRequest() *DistributionOrderSaveOpenRequest {
	return r._orderRequest
}

var poolTaobaoMiniappDistributionOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappDistributionOrderCreateRequest()
	},
}

// GetTaobaoMiniappDistributionOrderCreateRequest 从 sync.Pool 获取 TaobaoMiniappDistributionOrderCreateAPIRequest
func GetTaobaoMiniappDistributionOrderCreateAPIRequest() *TaobaoMiniappDistributionOrderCreateAPIRequest {
	return poolTaobaoMiniappDistributionOrderCreateAPIRequest.Get().(*TaobaoMiniappDistributionOrderCreateAPIRequest)
}

// ReleaseTaobaoMiniappDistributionOrderCreateAPIRequest 将 TaobaoMiniappDistributionOrderCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappDistributionOrderCreateAPIRequest(v *TaobaoMiniappDistributionOrderCreateAPIRequest) {
	v.Reset()
	poolTaobaoMiniappDistributionOrderCreateAPIRequest.Put(v)
}
