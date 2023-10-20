package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvpvlist 闲鱼已验货pv查询
// alibaba.idle.isv.pv.list
//
// 根据闲鱼渠道类目查询对应的品牌和型号清单，供服务商进行选择
func Alibabaidleisvpvlist(clt *core.SDKClient, req *idleisv.AlibabaidleisvpvlistAPIRequest, session string) (*idleisv.AlibabaidleisvpvlistAPIResponse, error) {
	var resp idleisv.AlibabaidleisvpvlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
