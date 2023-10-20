package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentitemadd 租赁商品发布
// alibaba.idle.rent.item.add
//
// 发布闲鱼租赁商品
func Alibabaidlerentitemadd(clt *core.SDKClient, req *idle.AlibabaidlerentitemaddAPIRequest, session string) (*idle.AlibabaidlerentitemaddAPIResponse, error) {
	var resp idle.AlibabaidlerentitemaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
