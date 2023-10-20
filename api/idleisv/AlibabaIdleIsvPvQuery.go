package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvpvquery 查询pv属性
// alibaba.idle.isv.pv.query
//
// 查询pv属性
func Alibabaidleisvpvquery(clt *core.SDKClient, req *idleisv.AlibabaidleisvpvqueryAPIRequest, session string) (*idleisv.AlibabaidleisvpvqueryAPIResponse, error) {
	var resp idleisv.AlibabaidleisvpvqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
