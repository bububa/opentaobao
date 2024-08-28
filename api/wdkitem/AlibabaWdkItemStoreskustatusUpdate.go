package wdkitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemStoreskustatusUpdate 修改门店商品状态
// alibaba.wdk.item.storeskustatus.update
//
// 五道口商品 修改门店商品状态
func AlibabaWdkItemStoreskustatusUpdate(ctx context.Context, clt *core.SDKClient, req *wdkitem.AlibabaWdkItemStoreskustatusUpdateAPIRequest, resp *wdkitem.AlibabaWdkItemStoreskustatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
