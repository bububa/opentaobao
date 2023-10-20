package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitemdiscountqueryitems 查询特价商品
// alibaba.hm.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
func Alibabahmmarketingitemdiscountqueryitems(clt *core.SDKClient, req *wdk.AlibabahmmarketingitemdiscountqueryitemsAPIRequest, session string) (*wdk.AlibabahmmarketingitemdiscountqueryitemsAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitemdiscountqueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
