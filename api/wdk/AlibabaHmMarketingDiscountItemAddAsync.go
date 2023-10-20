package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingdiscountitemaddasync 特价批量新增商品
// alibaba.hm.marketing.discount.item.add.async
//
// 新分组模型下新增商品
func Alibabahmmarketingdiscountitemaddasync(clt *core.SDKClient, req *wdk.AlibabahmmarketingdiscountitemaddasyncAPIRequest, session string) (*wdk.AlibabahmmarketingdiscountitemaddasyncAPIResponse, error) {
	var resp wdk.AlibabahmmarketingdiscountitemaddasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
