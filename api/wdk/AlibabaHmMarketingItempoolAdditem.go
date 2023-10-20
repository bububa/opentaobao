package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempooladditem 增加商品池里面的商品
// alibaba.hm.marketing.itempool.additem
//
// 增加商品池里面的商品
func Alibabahmmarketingitempooladditem(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempooladditemAPIRequest, session string) (*wdk.AlibabahmmarketingitempooladditemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempooladditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
