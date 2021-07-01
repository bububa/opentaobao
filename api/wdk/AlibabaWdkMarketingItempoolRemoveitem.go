package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkMarketingItempoolRemoveitem
移除商品池里面的商品
alibaba.wdk.marketing.itempool.removeitem

移除商品池里面的商品 */
func AlibabaWdkMarketingItempoolRemoveitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolRemoveitemAPIRequest, session string) (*wdk.AlibabaWdkMarketingItempoolRemoveitemAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingItempoolRemoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
