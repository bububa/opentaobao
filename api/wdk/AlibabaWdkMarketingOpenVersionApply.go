package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingopenversionapply 数据同步版本号申请
// alibaba.wdk.marketing.open.version.apply
//
// 数据同步版本号申请
func Alibabawdkmarketingopenversionapply(clt *core.SDKClient, req *wdk.AlibabawdkmarketingopenversionapplyAPIRequest, session string) (*wdk.AlibabawdkmarketingopenversionapplyAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingopenversionapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
