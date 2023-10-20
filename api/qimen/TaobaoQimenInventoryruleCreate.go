package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimeninventoryrulecreate 渠道间库存规则设置接口
// taobao.qimen.inventoryrule.create
//
// 渠道间库存规则设置
func Taobaoqimeninventoryrulecreate(clt *core.SDKClient, req *qimen.TaobaoqimeninventoryrulecreateAPIRequest, session string) (*qimen.TaobaoqimeninventoryrulecreateAPIResponse, error) {
	var resp qimen.TaobaoqimeninventoryrulecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
