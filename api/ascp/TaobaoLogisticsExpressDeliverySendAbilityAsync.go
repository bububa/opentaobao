package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsexpressdeliverysendabilityasync 快递送货上门能力同步/更新接口
// taobao.logistics.express.delivery.send.ability.async
//
// 快递送货上门能力同步/更新接口
func Taobaologisticsexpressdeliverysendabilityasync(clt *core.SDKClient, req *ascp.TaobaologisticsexpressdeliverysendabilityasyncAPIRequest, session string) (*ascp.TaobaologisticsexpressdeliverysendabilityasyncAPIResponse, error) {
	var resp ascp.TaobaologisticsexpressdeliverysendabilityasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
