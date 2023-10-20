package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolitemaddasync 商品池新增商品
// alibaba.hm.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
func Alibabahmmarketingitempoolitemaddasync(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolitemaddasyncAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolitemaddasyncAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolitemaddasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
