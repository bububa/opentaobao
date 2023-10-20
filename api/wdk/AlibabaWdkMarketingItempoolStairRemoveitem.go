package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolstairremoveitem 删除换购活动商品
// alibaba.wdk.marketing.itempool.stair.removeitem
//
// 删除换购商品
func Alibabawdkmarketingitempoolstairremoveitem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolstairremoveitemAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolstairremoveitemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolstairremoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
