package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitembuygiftqueryitems 查询买赠活动下的商品
// alibaba.hm.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
func Alibabahmmarketingitembuygiftqueryitems(clt *core.SDKClient, req *wdk.AlibabahmmarketingitembuygiftqueryitemsAPIRequest, session string) (*wdk.AlibabahmmarketingitembuygiftqueryitemsAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitembuygiftqueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
