package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolstairremoveitem 删除换购活动商品
// alibaba.hm.marketing.itempool.stair.removeitem
//
// 删除换购商品
func Alibabahmmarketingitempoolstairremoveitem(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolstairremoveitemAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolstairremoveitemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolstairremoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
