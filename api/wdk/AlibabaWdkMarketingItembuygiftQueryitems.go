package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitembuygiftqueryitems 查询买赠活动下的商品
// alibaba.wdk.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
func Alibabawdkmarketingitembuygiftqueryitems(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitembuygiftqueryitemsAPIRequest, session string) (*wdk.AlibabawdkmarketingitembuygiftqueryitemsAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitembuygiftqueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
