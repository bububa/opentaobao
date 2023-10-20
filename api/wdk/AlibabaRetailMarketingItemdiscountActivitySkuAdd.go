package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitemdiscountactivityskuadd 特价活动新增商品
// alibaba.retail.marketing.itemdiscount.activity.sku.add
//
// 新增或更新活动商品信息【同城零售】
func Alibabaretailmarketingitemdiscountactivityskuadd(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitemdiscountactivityskuaddAPIRequest, session string) (*wdk.AlibabaretailmarketingitemdiscountactivityskuaddAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitemdiscountactivityskuaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
