package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* AlibabaSscSupplyplatformServiceInventoryQuery
服务库存查询
alibaba.ssc.supplyplatform.service.inventory.query

查询服务库存。需要保存服务容量成功后，才能进行查询，参数中的provider信息（provider_id和provider_type）与alibaba.ssc.supplyplatform.servicecapacity.save接口中保持一致。 */
func AlibabaSscSupplyplatformServiceInventoryQuery(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServiceInventoryQueryAPIResponse, error) {
	var resp tmallservice.AlibabaSscSupplyplatformServiceInventoryQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
