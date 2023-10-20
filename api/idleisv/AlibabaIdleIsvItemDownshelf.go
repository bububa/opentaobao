package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvitemdownshelf 服务商闲鱼商品下架
// alibaba.idle.isv.item.downshelf
//
// 供外部服务商ISV进行闲鱼商品下架操作
func Alibabaidleisvitemdownshelf(clt *core.SDKClient, req *idleisv.AlibabaidleisvitemdownshelfAPIRequest, session string) (*idleisv.AlibabaidleisvitemdownshelfAPIResponse, error) {
	var resp idleisv.AlibabaidleisvitemdownshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
