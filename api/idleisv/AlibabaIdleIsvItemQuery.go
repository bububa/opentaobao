package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvitemquery 服务商闲鱼商品查询
// alibaba.idle.isv.item.query
//
// 服务商ISV闲鱼商品查询
func Alibabaidleisvitemquery(clt *core.SDKClient, req *idleisv.AlibabaidleisvitemqueryAPIRequest, session string) (*idleisv.AlibabaidleisvitemqueryAPIResponse, error) {
	var resp idleisv.AlibabaidleisvitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
