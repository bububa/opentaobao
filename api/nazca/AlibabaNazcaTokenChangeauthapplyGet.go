package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// Alibabanazcatokenchangeauthapplyget 根据token获取变更认证申请信息
// alibaba.nazca.token.changeauthapply.get
//
// 根据token获取变更认证申请信息
func Alibabanazcatokenchangeauthapplyget(clt *core.SDKClient, req *nazca.AlibabanazcatokenchangeauthapplygetAPIRequest, session string) (*nazca.AlibabanazcatokenchangeauthapplygetAPIResponse, error) {
	var resp nazca.AlibabanazcatokenchangeauthapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
