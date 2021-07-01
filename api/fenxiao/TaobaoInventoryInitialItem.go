package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoInventoryInitialItem
商品库存初始化
taobao.inventory.initial.item

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓商品初始化在各个仓中库存 */
func TaobaoInventoryInitialItem(clt *core.SDKClient, req *fenxiao.TaobaoInventoryInitialItemAPIRequest, session string) (*fenxiao.TaobaoInventoryInitialItemAPIResponse, error) {
	var resp fenxiao.TaobaoInventoryInitialItemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
