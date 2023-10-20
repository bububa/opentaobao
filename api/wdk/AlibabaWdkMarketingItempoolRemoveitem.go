package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolremoveitem 移除商品池里面的商品
// alibaba.wdk.marketing.itempool.removeitem
//
// 移除商品池里面的商品
func Alibabawdkmarketingitempoolremoveitem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolremoveitemAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolremoveitemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolremoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
