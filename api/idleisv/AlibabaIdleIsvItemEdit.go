package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvitemedit 服务商闲鱼商品编辑
// alibaba.idle.isv.item.edit
//
// 服务商ISV闲鱼商品编辑操作
func Alibabaidleisvitemedit(clt *core.SDKClient, req *idleisv.AlibabaidleisvitemeditAPIRequest, session string) (*idleisv.AlibabaidleisvitemeditAPIResponse, error) {
	var resp idleisv.AlibabaidleisvitemeditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
