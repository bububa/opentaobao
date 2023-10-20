package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeletevenue 大麦换验平台-第三方对外开放-场馆接口deleteVenue
// alibaba.damai.mev.open.deletevenue
//
// 开放接口，删除场馆
func AlibabaDamaiMevOpenDeletevenue(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletevenueAPIRequest, resp *damai.AlibabaDamaiMevOpenDeletevenueAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
