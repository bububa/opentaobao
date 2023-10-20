package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkCouponContractCreate 营销券合同创建接口
// alibaba.wdk.coupon.contract.create
//
// 营销券合同创建接口
func AlibabaWdkCouponContractCreate(clt *core.SDKClient, req *wdk.AlibabaWdkCouponContractCreateAPIRequest, session string) (*wdk.AlibabaWdkCouponContractCreateAPIResponse, error) {
	var resp wdk.AlibabaWdkCouponContractCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
