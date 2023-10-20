package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopendeleteperform 大麦换验平台-第三方对外开放-场次接口deletePerform
// alibaba.damai.mev.open.deleteperform
//
// deletePerform
func Alibabadamaimevopendeleteperform(clt *core.SDKClient, req *damai.AlibabadamaimevopendeleteperformAPIRequest, session string) (*damai.AlibabadamaimevopendeleteperformAPIResponse, error) {
	var resp damai.AlibabadamaimevopendeleteperformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
