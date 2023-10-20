package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushperform 大麦换验平台-第三方对外开放-场次接口pushPerform
// alibaba.damai.mev.open.pushperform
//
// pushPerform
func AlibabaDamaiMevOpenPushperform(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushperformAPIRequest, session string) (*damai.AlibabaDamaiMevOpenPushperformAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenPushperformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
