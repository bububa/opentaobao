package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopendeletevenue 大麦换验平台-第三方对外开放-场馆接口deleteVenue
// alibaba.damai.mev.open.deletevenue
//
// 开放接口，删除场馆
func Alibabadamaimevopendeletevenue(clt *core.SDKClient, req *damai.AlibabadamaimevopendeletevenueAPIRequest, session string) (*damai.AlibabadamaimevopendeletevenueAPIResponse, error) {
	var resp damai.AlibabadamaimevopendeletevenueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
