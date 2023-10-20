package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuChannelskuUpdate 更新渠道商品
// alibaba.wdk.sku.channelsku.update
//
// 批量更新渠道商品，商家通过Top接入
func AlibabaWdkSkuChannelskuUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkSkuChannelskuUpdateAPIRequest, resp *wdk.AlibabaWdkSkuChannelskuUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
