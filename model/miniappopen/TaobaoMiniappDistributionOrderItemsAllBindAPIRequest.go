package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionOrderItemsAllBindAPIRequest 小程序投放-基于投放计划绑定/解绑全店商品 API请求
// taobao.miniapp.distribution.order.items.all.bind
//
// 提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。该接口对应的能力是全店投放，表示店铺里的所有商品均会透出对应的投放计划（并且对于新上架的商品，系统也会每天凌晨进行一次补充投放，始终保持全店商品都有悬浮窗入口），该能力的使用，需要联系平台运营进行人工申请，申请通过以后方可调用成功。
type TaobaoMiniappDistributionOrderItemsAllBindAPIRequest struct {
	model.Params
	// 绑定/解绑的入参信息
	_allItemBindRequest *DistributionOrderBindTargetEntityOpenRequestV2
}

// NewTaobaoMiniappDistributionOrderItemsAllBindRequest 初始化TaobaoMiniappDistributionOrderItemsAllBindAPIRequest对象
func NewTaobaoMiniappDistributionOrderItemsAllBindRequest() *TaobaoMiniappDistributionOrderItemsAllBindAPIRequest {
	return &TaobaoMiniappDistributionOrderItemsAllBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionOrderItemsAllBindAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.order.items.all.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionOrderItemsAllBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappDistributionOrderItemsAllBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAllItemBindRequest is AllItemBindRequest Setter
// 绑定/解绑的入参信息
func (r *TaobaoMiniappDistributionOrderItemsAllBindAPIRequest) SetAllItemBindRequest(_allItemBindRequest *DistributionOrderBindTargetEntityOpenRequestV2) error {
	r._allItemBindRequest = _allItemBindRequest
	r.Set("all_item_bind_request", _allItemBindRequest)
	return nil
}

// GetAllItemBindRequest AllItemBindRequest Getter
func (r TaobaoMiniappDistributionOrderItemsAllBindAPIRequest) GetAllItemBindRequest() *DistributionOrderBindTargetEntityOpenRequestV2 {
	return r._allItemBindRequest
}
