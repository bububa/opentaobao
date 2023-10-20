package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolStairRemoveitem 删除换购活动商品
// alibaba.hm.marketing.itempool.stair.removeitem
//
// 删除换购商品
func AlibabaHmMarketingItempoolStairRemoveitem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolStairRemoveitemAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolStairRemoveitemAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolStairRemoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
