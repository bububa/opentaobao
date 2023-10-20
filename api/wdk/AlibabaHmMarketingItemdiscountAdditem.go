package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitemdiscountadditem 报名特价商品
// alibaba.hm.marketing.itemdiscount.additem
//
// 在商品特价活动中报名特价商品
func Alibabahmmarketingitemdiscountadditem(clt *core.SDKClient, req *wdk.AlibabahmmarketingitemdiscountadditemAPIRequest, session string) (*wdk.AlibabahmmarketingitemdiscountadditemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitemdiscountadditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
