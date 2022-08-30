package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolSkuQuery 查询商品池活动商品【同城零售】
// alibaba.retail.marketing.itempool.sku.query
//
// 查询商品池活动商品【同城零售】
func AlibabaRetailMarketingItempoolSkuQuery(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolSkuQueryAPIRequest, session string) (*wdk.AlibabaRetailMarketingItempoolSkuQueryAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingItempoolSkuQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
