package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolRemoveitem 移除商品池里面的商品
// alibaba.hm.marketing.itempool.removeitem
//
// 移除商品池里面的商品
func AlibabaHmMarketingItempoolRemoveitem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolRemoveitemAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolRemoveitemAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolRemoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
