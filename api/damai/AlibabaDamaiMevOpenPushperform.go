package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushperform 大麦换验平台-第三方对外开放-场次接口pushPerform
// alibaba.damai.mev.open.pushperform
//
// pushPerform
func AlibabaDamaiMevOpenPushperform(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushperformAPIRequest, resp *damai.AlibabaDamaiMevOpenPushperformAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
