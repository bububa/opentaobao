package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolActivityQuery 查询商品池活动【同城零售】
// alibaba.retail.marketing.itempool.activity.query
//
// 查询商品池活动【同城零售】
func AlibabaRetailMarketingItempoolActivityQuery(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivityQueryAPIRequest, session string) (*wdk.AlibabaRetailMarketingItempoolActivityQueryAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingItempoolActivityQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
