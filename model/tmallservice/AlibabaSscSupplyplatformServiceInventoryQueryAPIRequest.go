package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest
服务库存查询 API请求
alibaba.ssc.supplyplatform.service.inventory.query

查询服务库存。需要保存服务容量成功后，才能进行查询，参数中的provider信息（provider_id和provider_type）与alibaba.ssc.supplyplatform.servicecapacity.save接口中保持一致。 */
type AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest struct {
	model.Params
	// 查询开始日期。yyyy-MM-dd格式
	_startDay string
	// 查询结束日期。与start间隔不能超过31天。yyyy-MM-dd格式
	_endDay string
	// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
	_providerType string
	// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
	_providerId int64
}

// New
