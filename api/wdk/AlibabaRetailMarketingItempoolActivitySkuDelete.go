package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolActivitySkuDelete 删除商品池活动商品【同城零售】
// alibaba.retail.marketing.itempool.activity.sku.delete
//
// 删除商品池活动商品信息【同城零售】
func AlibabaRetailMarketingItempoolActivitySkuDelete(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest, session string) (*wdk.AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
