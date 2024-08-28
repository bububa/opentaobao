package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkopenCateorderPull 商户回传餐饮加工单状态
// alibaba.wdkopen.cateorder.pull
//
// 商户回传餐饮加工单状态
func AlibabaWdkopenCateorderPull(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkopenCateorderPullAPIRequest, resp *wdk.AlibabaWdkopenCateorderPullAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
