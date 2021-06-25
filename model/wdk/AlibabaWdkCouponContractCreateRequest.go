package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
营销券合同创建接口 APIRequest
alibaba.wdk.coupon.contract.create

营销券合同创建接口
*/
type AlibabaWdkCouponContractCreateRequest struct {
    model.Params

    // 调用入参
    createContractInstanceRequest   *CreateContractInstanceRequest 

}

func NewAlibabaWdkCouponContractCreateRequest() *AlibabaWdkCouponContractCreateRequest{
    return &AlibabaWdkCouponContractCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkCouponContractCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.contract.create"
}

func (r AlibabaWdkCouponContractCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkCouponContractCreateRequest) SetCreateContractInstanceRequest(createContractInstanceRequest *CreateContractInstanceRequest) error {
    r.createContractInstanceRequest = createContractInstanceRequest
    r.Set("create_contract_instance_request", createContractInstanceRequest)
    return nil
}

func (r AlibabaWdkCouponContractCreateRequest) GetCreateContractInstanceRequest() *CreateContractInstanceRequest {
    return r.createContractInstanceRequest
}

