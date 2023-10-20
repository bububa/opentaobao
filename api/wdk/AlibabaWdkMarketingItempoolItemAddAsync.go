package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolItemAddAsync 商品池新增商品
// alibaba.wdk.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
func AlibabaWdkMarketingItempoolItemAddAsync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest, session string) (*wdk.AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
