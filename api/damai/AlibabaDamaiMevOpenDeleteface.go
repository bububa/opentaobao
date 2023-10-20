package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeleteface 大麦换验平台-第三方对外开放-票面接口deleteFace
// alibaba.damai.mev.open.deleteface
//
// deleteFace
func AlibabaDamaiMevOpenDeleteface(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletefaceAPIRequest, resp *damai.AlibabaDamaiMevOpenDeletefaceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
