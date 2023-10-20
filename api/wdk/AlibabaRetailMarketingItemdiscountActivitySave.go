package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitemdiscountactivitysave 【同城零售】单品活动保存
// alibaba.retail.marketing.itemdiscount.activity.save
//
// 【同城零售】单品活动保存
func Alibabaretailmarketingitemdiscountactivitysave(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitemdiscountactivitysaveAPIRequest, session string) (*wdk.AlibabaretailmarketingitemdiscountactivitysaveAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitemdiscountactivitysaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
