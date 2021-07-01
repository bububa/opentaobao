package vms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoVmsServiceVehicleinfoUploadAPIRequest
新能源车--外部车辆信息回传 API请求
cainiao.vms.service.vehicleinfo.upload

新能源车--外部车辆信息回传 */
type CainiaoVmsServiceVehicleinfoUploadAPIRequest struct {
	model.Params
	// 设备号
	_deviceId string
	// 厂家标识
	_providerName string
	// 数据源标识
	_dataSource string
	// 协议版本标识
	_protocolVersion string
	// 上传的信息
	_data string
}

// New
