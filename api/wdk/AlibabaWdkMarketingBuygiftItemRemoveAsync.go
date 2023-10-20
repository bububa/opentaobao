package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingbuygiftitemremoveasync 批量删除买赠商品
// alibaba.wdk.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
func Alibabawdkmarketingbuygiftitemremoveasync(clt *core.SDKClient, req *wdk.AlibabawdkmarketingbuygiftitemremoveasyncAPIRequest, session string) (*wdk.AlibabawdkmarketingbuygiftitemremoveasyncAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingbuygiftitemremoveasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
