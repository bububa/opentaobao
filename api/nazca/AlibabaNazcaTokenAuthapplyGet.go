package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// AlibabaNazcaTokenAuthapplyGet 根据token获取认证申请信息
// alibaba.nazca.token.authapply.get
//
// 根据token获取认证申请信息
func AlibabaNazcaTokenAuthapplyGet(clt *core.SDKClient, req *nazca.AlibabaNazcaTokenAuthapplyGetAPIRequest, session string) (*nazca.AlibabaNazcaTokenAuthapplyGetAPIResponse, error) {
	var resp nazca.AlibabaNazcaTokenAuthapplyGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
