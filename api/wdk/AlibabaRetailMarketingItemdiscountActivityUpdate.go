package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitemdiscountactivityupdate 更新单品特价活动【同城零售】
// alibaba.retail.marketing.itemdiscount.activity.update
//
// 同城零售单品特价活动更新
func Alibabaretailmarketingitemdiscountactivityupdate(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitemdiscountactivityupdateAPIRequest, session string) (*wdk.AlibabaretailmarketingitemdiscountactivityupdateAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitemdiscountactivityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
