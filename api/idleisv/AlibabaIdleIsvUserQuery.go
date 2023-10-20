package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvuserquery 服务商ISV闲鱼用户信息查询
// alibaba.idle.isv.user.query
//
// 服务商ISV闲鱼用户信息查询
func Alibabaidleisvuserquery(clt *core.SDKClient, req *idleisv.AlibabaidleisvuserqueryAPIRequest, session string) (*idleisv.AlibabaidleisvuserqueryAPIResponse, error) {
	var resp idleisv.AlibabaidleisvuserqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
