package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// Alibabanazcatokenissuecertapplyget 根据token获取出证申请信息
// alibaba.nazca.token.issuecertapply.get
//
// 根据token获取出证申请信息
func Alibabanazcatokenissuecertapplyget(clt *core.SDKClient, req *nazca.AlibabanazcatokenissuecertapplygetAPIRequest, session string) (*nazca.AlibabanazcatokenissuecertapplygetAPIResponse, error) {
	var resp nazca.AlibabanazcatokenissuecertapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
