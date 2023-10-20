package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopendeletefloor 大麦换验平台-第三方对外开放-楼层接口deleteFloor
// alibaba.damai.mev.open.deletefloor
//
// deleteFloor
func Alibabadamaimevopendeletefloor(clt *core.SDKClient, req *damai.AlibabadamaimevopendeletefloorAPIRequest, session string) (*damai.AlibabadamaimevopendeletefloorAPIResponse, error) {
	var resp damai.AlibabadamaimevopendeletefloorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
