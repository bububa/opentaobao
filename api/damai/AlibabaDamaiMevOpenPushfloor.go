package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenpushfloor 大麦换验平台-第三方对外开放-楼层接口pushFloor
// alibaba.damai.mev.open.pushfloor
//
// pushFloor
func Alibabadamaimevopenpushfloor(clt *core.SDKClient, req *damai.AlibabadamaimevopenpushfloorAPIRequest, session string) (*damai.AlibabadamaimevopenpushfloorAPIResponse, error) {
	var resp damai.AlibabadamaimevopenpushfloorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
