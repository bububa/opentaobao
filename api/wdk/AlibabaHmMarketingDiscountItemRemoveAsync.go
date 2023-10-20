package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingdiscountitemremoveasync 特价批量移除商品
// alibaba.hm.marketing.discount.item.remove.async
//
// 特价批量移除商品
func Alibabahmmarketingdiscountitemremoveasync(clt *core.SDKClient, req *wdk.AlibabahmmarketingdiscountitemremoveasyncAPIRequest, session string) (*wdk.AlibabahmmarketingdiscountitemremoveasyncAPIResponse, error) {
	var resp wdk.AlibabahmmarketingdiscountitemremoveasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
