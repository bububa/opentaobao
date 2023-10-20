package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingbuygiftitemaddasync 批量发布买赠商品
// alibaba.hm.marketing.buygift.item.add.async
//
// 批量发布买赠商品
func Alibabahmmarketingbuygiftitemaddasync(clt *core.SDKClient, req *wdk.AlibabahmmarketingbuygiftitemaddasyncAPIRequest, session string) (*wdk.AlibabahmmarketingbuygiftitemaddasyncAPIResponse, error) {
	var resp wdk.AlibabahmmarketingbuygiftitemaddasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
