package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponContractCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.contract.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponContractCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CreateContractInstanceRequest Setter
// 调用入参
func (r *AlibabaWdkCouponContractCreateAPIRequest) SetCreateContractInstanceRequest(_createContractInstanceRequest *CreateContractInstanceRequest) error {
	r._createContractInstanceRequest = _createContractInstanceRequest
	r.Set("create_contract_instance_request", _createContractInstanceRequest)
	return nil
}

// Get CreateContractInstanceRequest Getter
func (r AlibabaWdkCouponContractCreateAPIRequest) GetCreateContractInstanceRequest() *CreateContractInstanceRequest {
	return r._createContractInstanceRequest
}
