package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolAdditem 增加商品池里面的商品
// alibaba.hm.marketing.itempool.additem
//
// 增加商品池里面的商品
func AlibabaHmMarketingItempoolAdditem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolAdditemAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolAdditemAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolAdditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
