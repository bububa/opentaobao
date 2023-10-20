package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitempoolactivitysave 【同城零售】保存商品池活动
// alibaba.retail.marketing.itempool.activity.save
//
// 同城零售商品池活动保存
func Alibabaretailmarketingitempoolactivitysave(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitempoolactivitysaveAPIRequest, session string) (*wdk.AlibabaretailmarketingitempoolactivitysaveAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitempoolactivitysaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
