package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionOrderGetAPIRequest 小程序投放-查询小程序投放计划信息 API请求
// taobao.miniapp.distribution.order.get
//
// 服务商可通过该API，读取自己开发的小程序对应的投放计划的相关信息
type TaobaoMiniappDistributionOrderGetAPIRequest struct {
	model.Params
	// 查询入参
	_orderIdRequest *DistributionOrderQueryByIdOpenRequest
}

// NewTaobaoMiniappDistributionOrderGetRequest 初始化TaobaoMiniappDistributionOrderGetAPIRequest对象
func NewTaobaoMiniappDistributionOrderGetRequest() *TaobaoMiniappDistributionOrderGetAPIRequest {
	return &TaobaoMiniappDistributionOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappDistributionOrderGetAPIRequest) Reset() {
	r._orderIdRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappDistributionOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderIdRequest is OrderIdRequest Setter
// 查询入参
func (r *TaobaoMiniappDistributionOrderGetAPIRequest) SetOrderIdRequest(_orderIdRequest *DistributionOrderQueryByIdOpenRequest) error {
	r._orderIdRequest = _orderIdRequest
	r.Set("order_id_request", _orderIdRequest)
	return nil
}

// GetOrderIdRequest OrderIdRequest Getter
func (r TaobaoMiniappDistributionOrderGetAPIRequest) GetOrderIdRequest() *DistributionOrderQueryByIdOpenRequest {
	return r._orderIdRequest
}

var poolTaobaoMiniappDistributionOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappDistributionOrderGetRequest()
	},
}

// GetTaobaoMiniappDistributionOrderGetRequest 从 sync.Pool 获取 TaobaoMiniappDistributionOrderGetAPIRequest
func GetTaobaoMiniappDistributionOrderGetAPIRequest() *TaobaoMiniappDistributionOrderGetAPIRequest {
	return poolTaobaoMiniappDistributionOrderGetAPIRequest.Get().(*TaobaoMiniappDistributionOrderGetAPIRequest)
}

// ReleaseTaobaoMiniappDistributionOrderGetAPIRequest 将 TaobaoMiniappDistributionOrderGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappDistributionOrderGetAPIRequest(v *TaobaoMiniappDistributionOrderGetAPIRequest) {
	v.Reset()
	poolTaobaoMiniappDistributionOrderGetAPIRequest.Put(v)
}
