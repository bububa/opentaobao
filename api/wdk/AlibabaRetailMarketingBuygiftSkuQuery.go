package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftSkuQuery 查询买赠活动商品【同城零售】
// alibaba.retail.marketing.buygift.sku.query
//
// 查询买赠活动商品【同城零售】
func AlibabaRetailMarketingBuygiftSkuQuery(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftSkuQueryAPIRequest, session string) (*wdk.AlibabaRetailMarketingBuygiftSkuQueryAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingBuygiftSkuQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
