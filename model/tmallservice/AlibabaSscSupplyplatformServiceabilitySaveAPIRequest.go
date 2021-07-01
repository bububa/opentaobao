package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSscSupplyplatformServiceabilitySaveAPIRequest
保存服务能力 API请求
alibaba.ssc.supplyplatform.serviceability.save

保存服务能力 */
type AlibabaSscSupplyplatformServiceabilitySaveAPIRequest struct {
	model.Params
	// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
	_providerType string
	// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
	_providerId int64
	// 目前包含三种。in_store 到店；at_home 上门；transmit_service 寄修。请根据实际支持的履约类型填写
	_fulfilTypeList []string
	// 服务sku，具体的sku列表可以从服务商工作台的类目树获取
	_serviceSkuCodeList []string
	// 菜鸟地址编码，各级地址均可（全国、省、市、区、街道），根据实际支持的地区填写。当支持的履约类型包含上门时，必填
	_areaCodeList []int64
}

// New
