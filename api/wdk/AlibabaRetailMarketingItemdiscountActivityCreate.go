package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitemdiscountactivitycreate 创建单品特价活动【同城零售】
// alibaba.retail.marketing.itemdiscount.activity.create
//
// 同城零售单品特价活动创建
func Alibabaretailmarketingitemdiscountactivitycreate(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitemdiscountactivitycreateAPIRequest, session string) (*wdk.AlibabaretailmarketingitemdiscountactivitycreateAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitemdiscountactivitycreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
