package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingbuygiftitemaddasync 批量发布买赠商品
// alibaba.wdk.marketing.buygift.item.add.async
//
// 批量发布买赠商品
func Alibabawdkmarketingbuygiftitemaddasync(clt *core.SDKClient, req *wdk.AlibabawdkmarketingbuygiftitemaddasyncAPIRequest, session string) (*wdk.AlibabawdkmarketingbuygiftitemaddasyncAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingbuygiftitemaddasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
