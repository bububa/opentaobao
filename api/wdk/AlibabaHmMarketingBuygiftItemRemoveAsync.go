package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingBuygiftItemRemoveAsync 批量删除买赠商品
// alibaba.hm.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
func AlibabaHmMarketingBuygiftItemRemoveAsync(clt *core.SDKClient, req *wdk.AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest, session string) (*wdk.AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
