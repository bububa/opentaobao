package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitempoolactivitydelete 删除商品池活动【同城零售】
// alibaba.retail.marketing.itempool.activity.delete
//
// 同城零售商品池活动删除
func Alibabaretailmarketingitempoolactivitydelete(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitempoolactivitydeleteAPIRequest, session string) (*wdk.AlibabaretailmarketingitempoolactivitydeleteAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitempoolactivitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
