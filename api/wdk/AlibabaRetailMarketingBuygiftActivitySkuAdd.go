package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaRetailMarketingBuygiftActivitySkuAdd
添加单品买赠活动商品
alibaba.retail.marketing.buygift.activity.sku.add

新增或更新单品买赠活动商品信息【同城零售】 */
func AlibabaRetailMarketingBuygiftActivitySkuAdd(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest, session string) (*wdk.AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
