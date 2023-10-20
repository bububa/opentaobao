package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitemdiscountremoveitem 移除报名的商品
// alibaba.wdk.marketing.itemdiscount.removeitem
//
// 将报名特价活动的商品从特价活动中移除
func Alibabawdkmarketingitemdiscountremoveitem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitemdiscountremoveitemAPIRequest, session string) (*wdk.AlibabawdkmarketingitemdiscountremoveitemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitemdiscountremoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
