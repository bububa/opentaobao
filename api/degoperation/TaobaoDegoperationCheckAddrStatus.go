package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// Taobaodegoperationcheckaddrstatus 地址
// taobao.degoperation.check.addr.status
//
// 激励
func Taobaodegoperationcheckaddrstatus(clt *core.SDKClient, req *degoperation.TaobaodegoperationcheckaddrstatusAPIRequest, session string) (*degoperation.TaobaodegoperationcheckaddrstatusAPIResponse, error) {
	var resp degoperation.TaobaodegoperationcheckaddrstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
