package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushstand 大麦换验平台-第三方对外开放-看台接口pushStand
// alibaba.damai.mev.open.pushstand
//
// pushStand
func AlibabaDamaiMevOpenPushstand(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushstandAPIRequest, session string) (*damai.AlibabaDamaiMevOpenPushstandAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenPushstandAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
