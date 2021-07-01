package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappDistributionOrderGetAPIRequest
小程序投放-查询小程序投放计划信息 API请求
taobao.miniapp.distribution.order.get

服务商可通过该API，读取自己开发的小程序对应的投放计划的相关信息 */
type TaobaoMiniappDistributionOrderGetAPIRequest struct {
	model.Params
	// 查询入参
	_orderIdRequest *DistributionOrderQueryByIdOpenRequest
}

// NewTaobaoMiniappDistributionOrderGetRequest 初始化TaobaoMiniappDistributionOrderGetAPIRequest对象
func NewTaobaoMiniappDistributionOrderGetRequest() *TaobaoMiniappDistributionOrderGetAPIRequest {
	return &TaobaoMiniappDistributionOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderIdRequest Setter
// 查询入参
func (r *TaobaoMiniappDistributionOrderGetAPIRequest) SetOrderIdRequest(_orderIdRequest *DistributionOrderQueryByIdOpenRequest) error {
	r._orderIdRequest = _orderIdRequest
	r.Set("order_id_request", _orderIdRequest)
	return nil
}

// Get OrderIdRequest Getter
func (r TaobaoMiniappDistributionOrderGetAPIRequest) GetOrderIdRequest() *DistributionOrderQueryByIdOpenRequest {
	return r._orderIdRequest
}
