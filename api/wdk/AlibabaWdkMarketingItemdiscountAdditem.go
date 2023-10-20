package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitemdiscountadditem 报名特价商品
// alibaba.wdk.marketing.itemdiscount.additem
//
// 在商品特价活动中报名特价商品
func Alibabawdkmarketingitemdiscountadditem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitemdiscountadditemAPIRequest, session string) (*wdk.AlibabawdkmarketingitemdiscountadditemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitemdiscountadditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
