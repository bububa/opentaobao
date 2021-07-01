package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponContractCreateAPIRequest
营销券合同创建接口 API请求
alibaba.wdk.coupon.contract.create

营销券合同创建接口 */
type AlibabaWdkCouponContractCreateAPIRequest struct {
	model.Params
	// 调用入参
	_createContractInstanceRequest *CreateContractInstanceRequest
}

// New
