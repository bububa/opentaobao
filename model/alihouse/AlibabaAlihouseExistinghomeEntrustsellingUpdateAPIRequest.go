package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest) Reset() {
	r._customerEntrustSellingReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.entrustselling.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeEntrustsellingUpdateRequest()
	},
}

// GetAlibabaAlihouseExistinghomeEntrustsellingUpdateRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest
func GetAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest() *AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest {
	return poolAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest.Get().(*AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest 将 AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest(v *AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest.Put(v)
}
