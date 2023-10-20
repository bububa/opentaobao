package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopendeletestand 大麦换验平台-第三方对外开放-看台接口deleteStand
// alibaba.damai.mev.open.deletestand
//
// deleteStand
func Alibabadamaimevopendeletestand(clt *core.SDKClient, req *damai.AlibabadamaimevopendeletestandAPIRequest, session string) (*damai.AlibabadamaimevopendeletestandAPIResponse, error) {
	var resp damai.AlibabadamaimevopendeletestandAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
