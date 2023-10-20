package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitembuygiftadditem 增加买赠活动商品。【注意，此接口暂不支持并发！】
// alibaba.hm.marketing.itembuygift.additem
//
// 增加买赠活动商品。【注意，此接口暂不支持并发！】
func Alibabahmmarketingitembuygiftadditem(clt *core.SDKClient, req *wdk.AlibabahmmarketingitembuygiftadditemAPIRequest, session string) (*wdk.AlibabahmmarketingitembuygiftadditemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitembuygiftadditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
