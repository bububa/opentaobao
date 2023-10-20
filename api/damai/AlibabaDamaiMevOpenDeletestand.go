package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeletestand 大麦换验平台-第三方对外开放-看台接口deleteStand
// alibaba.damai.mev.open.deletestand
//
// deleteStand
func AlibabaDamaiMevOpenDeletestand(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletestandAPIRequest, resp *damai.AlibabaDamaiMevOpenDeletestandAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
