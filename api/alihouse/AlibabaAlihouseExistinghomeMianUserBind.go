package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomemianuserbind 主账号入驻
// alibaba.alihouse.existinghome.mian.user.bind
//
// 主账号入驻
func Alibabaalihouseexistinghomemianuserbind(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomemianuserbindAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomemianuserbindAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomemianuserbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
