package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAxChannelSkuStatusUpdate 翱象渠道商品上下架接口
// alibaba.ax.channel.sku.status.update
//
// 翱象渠道商品上下架接口
func AlibabaAxChannelSkuStatusUpdate(clt *core.SDKClient, req *wdk.AlibabaAxChannelSkuStatusUpdateAPIRequest, resp *wdk.AlibabaAxChannelSkuStatusUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
