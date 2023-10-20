package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolStairAdditem 商品池阶梯商品添加
// alibaba.hm.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
func AlibabaHmMarketingItempoolStairAdditem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolStairAdditemAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolStairAdditemAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolStairAdditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
