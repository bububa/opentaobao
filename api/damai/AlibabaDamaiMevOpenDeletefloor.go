package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeletefloor 大麦换验平台-第三方对外开放-楼层接口deleteFloor
// alibaba.damai.mev.open.deletefloor
//
// deleteFloor
func AlibabaDamaiMevOpenDeletefloor(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletefloorAPIRequest, resp *damai.AlibabaDamaiMevOpenDeletefloorAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
