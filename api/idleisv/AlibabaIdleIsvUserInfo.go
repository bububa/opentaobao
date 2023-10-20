package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvuserinfo 闲鱼用户信息查询接口
// alibaba.idle.isv.user.info
//
// 闲鱼用户信息查询接口
func Alibabaidleisvuserinfo(clt *core.SDKClient, req *idleisv.AlibabaidleisvuserinfoAPIRequest, session string) (*idleisv.AlibabaidleisvuserinfoAPIResponse, error) {
	var resp idleisv.AlibabaidleisvuserinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
