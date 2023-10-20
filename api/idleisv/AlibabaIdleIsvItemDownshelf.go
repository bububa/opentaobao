package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvItemDownshelf 服务商闲鱼商品下架
// alibaba.idle.isv.item.downshelf
//
// 供外部服务商ISV进行闲鱼商品下架操作
func AlibabaIdleIsvItemDownshelf(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvItemDownshelfAPIRequest, session string) (*idleisv.AlibabaIdleIsvItemDownshelfAPIResponse, error) {
	var resp idleisv.AlibabaIdleIsvItemDownshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
