package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAxChannelSkuStatusUpdate 翱象渠道商品上下架接口
// alibaba.ax.channel.sku.status.update
//
// 翱象渠道商品上下架接口
func AlibabaAxChannelSkuStatusUpdate(clt *core.SDKClient, req *wdk.AlibabaAxChannelSkuStatusUpdateAPIRequest, session string) (*wdk.AlibabaAxChannelSkuStatusUpdateAPIResponse, error) {
	var resp wdk.AlibabaAxChannelSkuStatusUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
