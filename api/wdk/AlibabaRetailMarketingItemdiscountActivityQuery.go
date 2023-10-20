package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitemdiscountactivityquery 查询单品特价活动【同城零售】
// alibaba.retail.marketing.itemdiscount.activity.query
//
// 查询单品特价活动【同城零售】
func Alibabaretailmarketingitemdiscountactivityquery(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitemdiscountactivityqueryAPIRequest, session string) (*wdk.AlibabaretailmarketingitemdiscountactivityqueryAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitemdiscountactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
