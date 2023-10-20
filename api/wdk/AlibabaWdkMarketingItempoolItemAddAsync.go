package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolitemaddasync 商品池新增商品
// alibaba.wdk.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
func Alibabawdkmarketingitempoolitemaddasync(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolitemaddasyncAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolitemaddasyncAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolitemaddasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
