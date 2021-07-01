package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSscSupplyplatformServicecapacitySaveAPIRequest
保存服务容量 API请求
alibaba.ssc.supplyplatform.servicecapacity.save

保存服务容量 */
type AlibabaSscSupplyplatformServicecapacitySaveAPIRequest struct {
	model.Params
	// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
	_providerType string
	// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
	_providerId int64
	// 目前支持两种。daily 每天容量相同；customize 定制容量，每天都不同
	_mode string
	// 目前支持两种。day 表示支持的时间粒度为天；hour 时间粒度为小时
	_timeInterval string
	// 容量数据。根据mode和time interval有不同的格式。具体格式见业务对接文档。
	_capacityData string
}

// New
