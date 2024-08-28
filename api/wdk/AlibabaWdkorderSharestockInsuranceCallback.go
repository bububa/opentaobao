package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkorderSharestockInsuranceCallback 共享库存订单投保后回传保单号
// alibaba.wdkorder.sharestock.insurance.callback
//
// 共享库存订单投保消息获取
func AlibabaWdkorderSharestockInsuranceCallback(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockInsuranceCallbackAPIRequest, resp *wdk.AlibabaWdkorderSharestockInsuranceCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
