package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitembuygiftadditem 增加买赠活动商品。【注意，此接口暂不支持并发！】
// alibaba.wdk.marketing.itembuygift.additem
//
// 增加买赠活动商品。【注意，此接口暂不支持并发！】
func Alibabawdkmarketingitembuygiftadditem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitembuygiftadditemAPIRequest, session string) (*wdk.AlibabawdkmarketingitembuygiftadditemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitembuygiftadditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
