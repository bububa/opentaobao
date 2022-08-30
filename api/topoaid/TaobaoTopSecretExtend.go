package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoTopSecretExtend 虚拟号延期
// taobao.top.secret.extend
//
// 虚拟号延期
func TaobaoTopSecretExtend(clt *core.SDKClient, req *topoaid.TaobaoTopSecretExtendAPIRequest, session string) (*topoaid.TaobaoTopSecretExtendAPIResponse, error) {
	var resp topoaid.TaobaoTopSecretExtendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
