package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuChannelskuAdd 新增渠道商品
// alibaba.wdk.sku.channelsku.add
//
// 盒马帮1期新增渠道商品
func AlibabaWdkSkuChannelskuAdd(clt *core.SDKClient, req *wdk.AlibabaWdkSkuChannelskuAddAPIRequest, resp *wdk.AlibabaWdkSkuChannelskuAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
