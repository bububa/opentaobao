package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceInventoryEdit 编辑服务库存
// alibaba.ssc.supplyplatform.service.inventory.edit
//
// 实时编辑服务库存。只支持增加或减少库存，不支持设置绝对库存值。
// 需要自己处理好幂等逻辑。
// 要先查询当前库存值，并基于返回结果做编辑操作。
// 参考alibaba.ssc.supplyplatform.service.inventory.query和alibaba.ssc.supplyplatform.servicecapacity.save
func AlibabaSscSupplyplatformServiceInventoryEdit(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceInventoryEditAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceInventoryEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
