package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventorySynchronize 库存状态同步接口
// taobao.qimen.inventory.synchronize
//
// ERP通过该接口同步指定商品的库存信息
func TaobaoQimenInventorySynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenInventorySynchronizeAPIRequest, session string) (*qimen.TaobaoQimenInventorySynchronizeAPIResponse, error) {
	var resp qimen.TaobaoQimenInventorySynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
