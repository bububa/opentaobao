package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeleteFaceelement 大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement
// alibaba.damai.mev.open.delete.faceelement
//
// deleteFaceElement
func AlibabaDamaiMevOpenDeleteFaceelement(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeleteFaceelementAPIRequest, resp *damai.AlibabaDamaiMevOpenDeleteFaceelementAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
