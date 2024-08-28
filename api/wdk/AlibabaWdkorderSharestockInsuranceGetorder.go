package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkorderSharestockInsuranceGetorder 共享库存订单投保消息获取
// alibaba.wdkorder.sharestock.insurance.getorder
//
// 共享库存订单投保消息获取
func AlibabaWdkorderSharestockInsuranceGetorder(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockInsuranceGetorderAPIRequest, resp *wdk.AlibabaWdkorderSharestockInsuranceGetorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
