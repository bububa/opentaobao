package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponContractCreateAPIRequest 营销券合同创建接口 API请求
// alibaba.wdk.coupon.contract.create
//
// 营销券合同创建接口
type AlibabaWdkCouponContractCreateAPIRequest struct {
	model.Params
	// 调用入参
	_createContractInstanceRequest *CreateContractInstanceRequest
}

// NewAlibabaWdkCouponContractCreateRequest 初始化AlibabaWdkCouponContractCreateAPIRequest对象
func NewAlibabaWdkCouponContractCreateRequest() *AlibabaWdkCouponContractCreateAPIRequest {
	return &AlibabaWdkCouponContractCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkCouponContractCreateAPIRequest) Reset() {
	r._createContractInstanceRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponContractCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.contract.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponContractCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkCouponContractCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateContractInstanceRequest is CreateContractInstanceRequest Setter
// 调用入参
func (r *AlibabaWdkCouponContractCreateAPIRequest) SetCreateContractInstanceRequest(_createContractInstanceRequest *CreateContractInstanceRequest) error {
	r._createContractInstanceRequest = _createContractInstanceRequest
	r.Set("create_contract_instance_request", _createContractInstanceRequest)
	return nil
}

// GetCreateContractInstanceRequest CreateContractInstanceRequest Getter
func (r AlibabaWdkCouponContractCreateAPIRequest) GetCreateContractInstanceRequest() *CreateContractInstanceRequest {
	return r._createContractInstanceRequest
}

var poolAlibabaWdkCouponContractCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkCouponContractCreateRequest()
	},
}

// GetAlibabaWdkCouponContractCreateRequest 从 sync.Pool 获取 AlibabaWdkCouponContractCreateAPIRequest
func GetAlibabaWdkCouponContractCreateAPIRequest() *AlibabaWdkCouponContractCreateAPIRequest {
	return poolAlibabaWdkCouponContractCreateAPIRequest.Get().(*AlibabaWdkCouponContractCreateAPIRequest)
}

// ReleaseAlibabaWdkCouponContractCreateAPIRequest 将 AlibabaWdkCouponContractCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkCouponContractCreateAPIRequest(v *AlibabaWdkCouponContractCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkCouponContractCreateAPIRequest.Put(v)
}
