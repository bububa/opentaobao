package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressDeliverySendAbilityAsync 快递送货上门能力同步/更新接口
// taobao.logistics.express.delivery.send.ability.async
//
// 快递送货上门能力同步/更新接口
func TaobaoLogisticsExpressDeliverySendAbilityAsync(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest, session string) (*ascp.TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse, error) {
	var resp ascp.TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
