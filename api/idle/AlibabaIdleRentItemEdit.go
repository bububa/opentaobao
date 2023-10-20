package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentitemedit 租赁商品编辑
// alibaba.idle.rent.item.edit
//
// 发布闲鱼租赁商品
func Alibabaidlerentitemedit(clt *core.SDKClient, req *idle.AlibabaidlerentitemeditAPIRequest, session string) (*idle.AlibabaidlerentitemeditAPIResponse, error) {
	var resp idle.AlibabaidlerentitemeditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
