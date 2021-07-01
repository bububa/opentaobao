package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

/* AlibabaDamaiMevOpenDeleteFaceelement
大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement
alibaba.damai.mev.open.delete.faceelement

deleteFaceElement */
func AlibabaDamaiMevOpenDeleteFaceelement(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeleteFaceelementAPIRequest, session string) (*damai.AlibabaDamaiMevOpenDeleteFaceelementAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenDeleteFaceelementAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
