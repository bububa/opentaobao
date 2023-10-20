package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingdiscountitemaddasync 特价批量新增商品
// alibaba.wdk.marketing.discount.item.add.async
//
// 新分组模型下新增商品
func Alibabawdkmarketingdiscountitemaddasync(clt *core.SDKClient, req *wdk.AlibabawdkmarketingdiscountitemaddasyncAPIRequest, session string) (*wdk.AlibabawdkmarketingdiscountitemaddasyncAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingdiscountitemaddasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
