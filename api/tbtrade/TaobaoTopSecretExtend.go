package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTopSecretExtend 虚拟号延期
// taobao.top.secret.extend
//
// 虚拟号延期
func TaobaoTopSecretExtend(clt *core.SDKClient, req *tbtrade.TaobaoTopSecretExtendAPIRequest, session string) (*tbtrade.TaobaoTopSecretExtendAPIResponse, error) {
	var resp tbtrade.TaobaoTopSecretExtendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
