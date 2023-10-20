package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempooladditem 增加商品池里面的商品
// alibaba.wdk.marketing.itempool.additem
//
// 增加商品池里面的商品
func Alibabawdkmarketingitempooladditem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempooladditemAPIRequest, session string) (*wdk.AlibabawdkmarketingitempooladditemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempooladditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
