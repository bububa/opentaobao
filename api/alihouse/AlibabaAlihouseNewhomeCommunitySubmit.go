package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCommunitySubmit 提交小区信息
// alibaba.alihouse.newhome.community.submit
//
// 提交小区信息
func AlibabaAlihouseNewhomeCommunitySubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCommunitySubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeCommunitySubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
