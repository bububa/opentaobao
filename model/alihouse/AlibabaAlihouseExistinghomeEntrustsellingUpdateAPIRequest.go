package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest 管家状态及房源信息接口 API请求
// alibaba.alihouse.existinghome.entrustselling.update
//
// 管家状态及房源信息接口
type AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest struct {
	model.Params
	// 参数
	_customerEntrustSellingReq *CustomerEntrustSellingReq
}

// NewAlibabaAlihouseExistinghomeEntrustsellingUpdateRequest 初始化AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest对象
func NewAlibabaAlihouseExistinghomeEntrustsellingUpdateRequest() *AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest {
	return &AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.entrustselling.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCustomerEntrustSellingReq is CustomerEntrustSellingReq Setter
// 参数
func (r *AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest) SetCustomerEntrustSellingReq(_customerEntrustSellingReq *CustomerEntrustSellingReq) error {
	r._customerEntrustSellingReq = _customerEntrustSellingReq
	r.Set("customer_entrust_selling_req", _customerEntrustSellingReq)
	return nil
}

// GetCustomerEntrustSellingReq CustomerEntrustSellingReq Getter
func (r AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest) GetCustomerEntrustSellingReq() *CustomerEntrustSellingReq {
	return r._customerEntrustSellingReq
}
