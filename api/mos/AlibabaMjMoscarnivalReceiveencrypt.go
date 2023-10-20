package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMoscarnivalReceiveencrypt 根据加密手机号领券
// alibaba.mj.moscarnival.receiveencrypt
//
// 根据加密手机号领券
func AlibabaMjMoscarnivalReceiveencrypt(clt *core.SDKClient, req *mos.AlibabaMjMoscarnivalReceiveencryptAPIRequest, resp *mos.AlibabaMjMoscarnivalReceiveencryptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
