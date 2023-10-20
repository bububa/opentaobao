package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenpushperform 大麦换验平台-第三方对外开放-场次接口pushPerform
// alibaba.damai.mev.open.pushperform
//
// pushPerform
func Alibabadamaimevopenpushperform(clt *core.SDKClient, req *damai.AlibabadamaimevopenpushperformAPIRequest, session string) (*damai.AlibabadamaimevopenpushperformAPIResponse, error) {
	var resp damai.AlibabadamaimevopenpushperformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
