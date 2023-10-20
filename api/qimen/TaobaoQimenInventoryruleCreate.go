package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventoryruleCreate 渠道间库存规则设置接口
// taobao.qimen.inventoryrule.create
//
// 渠道间库存规则设置
func TaobaoQimenInventoryruleCreate(clt *core.SDKClient, req *qimen.TaobaoQimenInventoryruleCreateAPIRequest, resp *qimen.TaobaoQimenInventoryruleCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
