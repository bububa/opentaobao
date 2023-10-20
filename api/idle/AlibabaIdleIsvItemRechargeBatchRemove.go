package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleisvitemrechargebatchremove 闲鱼商品直充功能移除
// alibaba.idle.isv.item.recharge.batch.remove
//
// 闲鱼商品直充功能移除
func Alibabaidleisvitemrechargebatchremove(clt *core.SDKClient, req *idle.AlibabaidleisvitemrechargebatchremoveAPIRequest, session string) (*idle.AlibabaidleisvitemrechargebatchremoveAPIResponse, error) {
	var resp idle.AlibabaidleisvitemrechargebatchremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
