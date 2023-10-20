package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCommunityLine 小区上下架
// alibaba.alihouse.newhome.community.line
//
// 小区上下架
func AlibabaAlihouseNewhomeCommunityLine(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCommunityLineAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeCommunityLineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
