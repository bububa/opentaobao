package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitempoolactivitycreate 创建商品池活动【同城零售】
// alibaba.retail.marketing.itempool.activity.create
//
// 同城零售商品池活动创建
func Alibabaretailmarketingitempoolactivitycreate(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitempoolactivitycreateAPIRequest, session string) (*wdk.AlibabaretailmarketingitempoolactivitycreateAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitempoolactivitycreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
