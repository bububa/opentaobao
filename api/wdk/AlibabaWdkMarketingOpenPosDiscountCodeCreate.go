package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingOpenPosDiscountCodeCreate pos一物一码创建
// alibaba.wdk.marketing.open.pos.discount.code.create
//
// pos一物一码创建
func AlibabaWdkMarketingOpenPosDiscountCodeCreate(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest, session string) (*wdk.AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
