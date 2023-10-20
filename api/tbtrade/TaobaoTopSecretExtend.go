package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaotopsecretextend 虚拟号延期
// taobao.top.secret.extend
//
// 虚拟号延期
func Taobaotopsecretextend(clt *core.SDKClient, req *tbtrade.TaobaotopsecretextendAPIRequest, session string) (*tbtrade.TaobaotopsecretextendAPIResponse, error) {
	var resp tbtrade.TaobaotopsecretextendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
