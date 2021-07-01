package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest
服务容量删除 API请求
alibaba.ssc.supplyplatform.servicecapacity.delete

服务容量删除 */
type AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest struct {
	model.Params
	// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
	_providerType string
	// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
	_providerId int64
}

// New
