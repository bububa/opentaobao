package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
营销券合同创建接口 API请求
alibaba.wdk.coupon.contract.create

营销券合同创建接口
*/
type AlibabaWdkCouponContractCreateRequest struct {
    model.Params
    // 调用入参
    _createContractInstanceRequest   *CreateContractInstanceRequest
}

// 初始化AlibabaWdkCouponContractCreateRequest对象
func NewAlibabaWdkCouponContractCreateRequest() *AlibabaWdkCouponContractCreateRequest{
    return &AlibabaWdkCouponContractCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponContractCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.contract.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponContractCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateContractInstanceRequest Setter
// 调用入参
func (r *AlibabaWdkCouponContractCreateRequest) SetCreateContractInstanceRequest(_createContractInstanceRequest *CreateContractInstanceRequest) error {
    r._createContractInstanceRequest = _createContractInstanceRequest
    r.Set("create_contract_instance_request", _createContractInstanceRequest)
    return nil
}

// CreateContractInstanceRequest Getter
func (r AlibabaWdkCouponContractCreateRequest) GetCreateContractInstanceRequest() *CreateContractInstanceRequest {
    return r._createContractInstanceRequest
}
