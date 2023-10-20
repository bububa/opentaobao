package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftRemoveitem 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
// alibaba.hm.marketing.itembuygift.removeitem
//
// 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
func AlibabaHmMarketingItembuygiftRemoveitem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftRemoveitemAPIRequest, session string) (*wdk.AlibabaHmMarketingItembuygiftRemoveitemAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItembuygiftRemoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
