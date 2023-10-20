package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushfaceelement 大麦换验平台-第三方对外开放-票面元素接口pushFaceElement
// alibaba.damai.mev.open.pushfaceelement
//
// pushFaceElement
func AlibabaDamaiMevOpenPushfaceelement(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushfaceelementAPIRequest, resp *damai.AlibabaDamaiMevOpenPushfaceelementAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
