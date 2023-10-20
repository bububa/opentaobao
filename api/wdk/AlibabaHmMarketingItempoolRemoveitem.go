package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolremoveitem 移除商品池里面的商品
// alibaba.hm.marketing.itempool.removeitem
//
// 移除商品池里面的商品
func Alibabahmmarketingitempoolremoveitem(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolremoveitemAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolremoveitemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolremoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
