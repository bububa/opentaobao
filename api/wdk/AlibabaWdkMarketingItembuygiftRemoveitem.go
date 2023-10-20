package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitembuygiftremoveitem 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
// alibaba.wdk.marketing.itembuygift.removeitem
//
// 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
func Alibabawdkmarketingitembuygiftremoveitem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitembuygiftremoveitemAPIRequest, session string) (*wdk.AlibabawdkmarketingitembuygiftremoveitemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitembuygiftremoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
