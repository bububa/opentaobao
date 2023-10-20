package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitemdiscountremoveitem 移除报名的商品
// alibaba.hm.marketing.itemdiscount.removeitem
//
// 将报名特价活动的商品从特价活动中移除
func Alibabahmmarketingitemdiscountremoveitem(clt *core.SDKClient, req *wdk.AlibabahmmarketingitemdiscountremoveitemAPIRequest, session string) (*wdk.AlibabahmmarketingitemdiscountremoveitemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitemdiscountremoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
