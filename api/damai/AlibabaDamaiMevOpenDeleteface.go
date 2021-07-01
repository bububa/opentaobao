package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

/* AlibabaDamaiMevOpenDeleteface
大麦换验平台-第三方对外开放-票面接口deleteFace
alibaba.damai.mev.open.deleteface

deleteFace */
func AlibabaDamaiMevOpenDeleteface(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletefaceAPIRequest, session string) (*damai.AlibabaDamaiMevOpenDeletefaceAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenDeletefaceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
