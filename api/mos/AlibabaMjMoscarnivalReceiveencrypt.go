package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjmoscarnivalreceiveencrypt 根据加密手机号领券
// alibaba.mj.moscarnival.receiveencrypt
//
// 根据加密手机号领券
func Alibabamjmoscarnivalreceiveencrypt(clt *core.SDKClient, req *mos.AlibabamjmoscarnivalreceiveencryptAPIRequest, session string) (*mos.AlibabamjmoscarnivalreceiveencryptAPIResponse, error) {
	var resp mos.AlibabamjmoscarnivalreceiveencryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
