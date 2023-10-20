package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingopenversioncount 版本数量查询
// alibaba.wdk.marketing.open.version.count
//
// 版本数量查询
func Alibabawdkmarketingopenversioncount(clt *core.SDKClient, req *wdk.AlibabawdkmarketingopenversioncountAPIRequest, session string) (*wdk.AlibabawdkmarketingopenversioncountAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingopenversioncountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
