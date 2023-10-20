package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleisvitemrechargeedit 闲鱼商品直充功能编辑
// alibaba.idle.isv.item.recharge.edit
//
// 闲鱼商品直充功能编辑
func Alibabaidleisvitemrechargeedit(clt *core.SDKClient, req *idle.AlibabaidleisvitemrechargeeditAPIRequest, session string) (*idle.AlibabaidleisvitemrechargeeditAPIResponse, error) {
	var resp idle.AlibabaidleisvitemrechargeeditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
