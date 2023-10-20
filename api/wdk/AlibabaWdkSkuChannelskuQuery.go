package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuChannelskuQuery 查询渠道商品
// alibaba.wdk.sku.channelsku.query
//
// 查询渠道商品
func AlibabaWdkSkuChannelskuQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuChannelskuQueryAPIRequest, resp *wdk.AlibabaWdkSkuChannelskuQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
