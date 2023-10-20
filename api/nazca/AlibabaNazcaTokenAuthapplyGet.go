package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// Alibabanazcatokenauthapplyget 根据token获取认证申请信息
// alibaba.nazca.token.authapply.get
//
// 根据token获取认证申请信息
func Alibabanazcatokenauthapplyget(clt *core.SDKClient, req *nazca.AlibabanazcatokenauthapplygetAPIRequest, session string) (*nazca.AlibabanazcatokenauthapplygetAPIResponse, error) {
	var resp nazca.AlibabanazcatokenauthapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
