package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

/* AlibabaDamaiMevOpenPushfloor
大麦换验平台-第三方对外开放-楼层接口pushFloor
alibaba.damai.mev.open.pushfloor

pushFloor */
func AlibabaDamaiMevOpenPushfloor(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushfloorAPIRequest, session string) (*damai.AlibabaDamaiMevOpenPushfloorAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenPushfloorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
