package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitemdiscountqueryitems 查询特价商品
// alibaba.wdk.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
func Alibabawdkmarketingitemdiscountqueryitems(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitemdiscountqueryitemsAPIRequest, session string) (*wdk.AlibabawdkmarketingitemdiscountqueryitemsAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitemdiscountqueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
