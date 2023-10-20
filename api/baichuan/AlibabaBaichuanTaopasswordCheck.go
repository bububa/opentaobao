package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Alibababaichuantaopasswordcheck 淘口令检查
// alibaba.baichuan.taopassword.check
//
// 检查当前文本是否为淘口令
func Alibababaichuantaopasswordcheck(clt *core.SDKClient, req *baichuan.AlibababaichuantaopasswordcheckAPIRequest, session string) (*baichuan.AlibababaichuantaopasswordcheckAPIResponse, error) {
	var resp baichuan.AlibababaichuantaopasswordcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
