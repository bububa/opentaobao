package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitempoolactivityquery 查询商品池活动【同城零售】
// alibaba.retail.marketing.itempool.activity.query
//
// 查询商品池活动【同城零售】
func Alibabaretailmarketingitempoolactivityquery(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitempoolactivityqueryAPIRequest, session string) (*wdk.AlibabaretailmarketingitempoolactivityqueryAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitempoolactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
