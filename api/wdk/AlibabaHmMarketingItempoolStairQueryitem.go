package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolStairQueryitem 换购商品查询
// alibaba.hm.marketing.itempool.stair.queryitem
//
// 换购商品查询
func AlibabaHmMarketingItempoolStairQueryitem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolStairQueryitemAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolStairQueryitemAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolStairQueryitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
