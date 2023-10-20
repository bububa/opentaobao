package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolItemRemoveAsync 商品池删除商品
// alibaba.hm.marketing.itempool.item.remove.async
//
// 新模型下删除商品
func AlibabaHmMarketingItempoolItemRemoveAsync(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolItemRemoveAsyncAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
