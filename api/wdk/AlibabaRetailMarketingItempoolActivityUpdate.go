package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitempoolactivityupdate 更新商品池活动【同城零售】
// alibaba.retail.marketing.itempool.activity.update
//
// 同城零售商品池活动更新
func Alibabaretailmarketingitempoolactivityupdate(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitempoolactivityupdateAPIRequest, session string) (*wdk.AlibabaretailmarketingitempoolactivityupdateAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitempoolactivityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
