package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeleteperform 大麦换验平台-第三方对外开放-场次接口deletePerform
// alibaba.damai.mev.open.deleteperform
//
// deletePerform
func AlibabaDamaiMevOpenDeleteperform(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeleteperformAPIRequest, resp *damai.AlibabaDamaiMevOpenDeleteperformAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
