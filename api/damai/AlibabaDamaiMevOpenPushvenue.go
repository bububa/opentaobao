package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushvenue 大麦换验平台-第三方对外开放-场馆接口pushVenue
// alibaba.damai.mev.open.pushvenue
//
// 开放接口推送场馆
func AlibabaDamaiMevOpenPushvenue(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushvenueAPIRequest, session string) (*damai.AlibabaDamaiMevOpenPushvenueAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenPushvenueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
