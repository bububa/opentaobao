package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushface 大麦换验平台-第三方对外开放-票面接口pushFace
// alibaba.damai.mev.open.pushface
//
// pushFace
func AlibabaDamaiMevOpenPushface(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushfaceAPIRequest, resp *damai.AlibabaDamaiMevOpenPushfaceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
