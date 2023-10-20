package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingdiscountitemremoveasync 特价批量移除商品
// alibaba.wdk.marketing.discount.item.remove.async
//
// 特价批量移除商品
func Alibabawdkmarketingdiscountitemremoveasync(clt *core.SDKClient, req *wdk.AlibabawdkmarketingdiscountitemremoveasyncAPIRequest, session string) (*wdk.AlibabawdkmarketingdiscountitemremoveasyncAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingdiscountitemremoveasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
