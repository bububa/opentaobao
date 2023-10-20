package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventoryruleCreate 渠道间库存规则设置接口
// taobao.qimen.inventoryrule.create
//
// 渠道间库存规则设置
func TaobaoQimenInventoryruleCreate(clt *core.SDKClient, req *qimen.TaobaoQimenInventoryruleCreateAPIRequest, session string) (*qimen.TaobaoQimenInventoryruleCreateAPIResponse, error) {
	var resp qimen.TaobaoQimenInventoryruleCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
