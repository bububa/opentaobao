package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenpushvenue 大麦换验平台-第三方对外开放-场馆接口pushVenue
// alibaba.damai.mev.open.pushvenue
//
// 开放接口推送场馆
func Alibabadamaimevopenpushvenue(clt *core.SDKClient, req *damai.AlibabadamaimevopenpushvenueAPIRequest, session string) (*damai.AlibabadamaimevopenpushvenueAPIResponse, error) {
	var resp damai.AlibabadamaimevopenpushvenueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
