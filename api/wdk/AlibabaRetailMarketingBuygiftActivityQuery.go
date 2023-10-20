package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingbuygiftactivityquery 查询单品买赠活动【同城零售】
// alibaba.retail.marketing.buygift.activity.query
//
// 查询单品买赠活动【同城零售】
func Alibabaretailmarketingbuygiftactivityquery(clt *core.SDKClient, req *wdk.AlibabaretailmarketingbuygiftactivityqueryAPIRequest, session string) (*wdk.AlibabaretailmarketingbuygiftactivityqueryAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingbuygiftactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
