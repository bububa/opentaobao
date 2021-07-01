package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

/* AlibabaDamaiMevOpenDeletefloor
大麦换验平台-第三方对外开放-楼层接口deleteFloor
alibaba.damai.mev.open.deletefloor

deleteFloor */
func AlibabaDamaiMevOpenDeletefloor(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletefloorAPIRequest, session string) (*damai.AlibabaDamaiMevOpenDeletefloorAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenDeletefloorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
