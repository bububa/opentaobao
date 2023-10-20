package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest 管家状态及房源信息接口 API请求
// alibaba.alihouse.existinghome.entrustselling.update
//
// 管家状态及房源信息接口
type AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest struct {
	model.Params
	// 参数
	_customerEntrustSellingReq *CustomerEntrustSellingReq
}

// NewAlibabaalihouseexistinghomeentrustsellingupdateRequest 初始化AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest对象
func NewAlibabaalihouseexistinghomeentrustsellingupdateRequest() *AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest {
	return &AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.entrustselling.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCustomerEntrustSellingReq is CustomerEntrustSellingReq Setter
// 参数
func (r *AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest) SetCustomerEntrustSellingReq(_customerEntrustSellingReq *CustomerEntrustSellingReq) error {
	r._customerEntrustSellingReq = _customerEntrustSellingReq
	r.Set("customer_entrust_selling_req", _customerEntrustSellingReq)
	return nil
}

// GetCustomerEntrustSellingReq CustomerEntrustSellingReq Getter
func (r AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest) GetCustomerEntrustSellingReq() *CustomerEntrustSellingReq {
	return r._customerEntrustSellingReq
}
