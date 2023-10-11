package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolItemAddAsync 商品池新增商品
// alibaba.hm.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
func AlibabaHmMarketingItempoolItemAddAsync(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolItemAddAsyncAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolItemAddAsyncAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolItemAddAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
