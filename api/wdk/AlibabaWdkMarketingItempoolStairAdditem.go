package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolStairAdditem 商品池阶梯商品添加
// alibaba.wdk.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
func AlibabaWdkMarketingItempoolStairAdditem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolStairAdditemAPIRequest, session string) (*wdk.AlibabaWdkMarketingItempoolStairAdditemAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingItempoolStairAdditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
