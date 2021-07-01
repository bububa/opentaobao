package ott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttAlicbFacadeserviceGetdataAPIRequest
影视SDK获取设备能力值 API请求
youku.ott.alicb.facadeservice.getdata

影视SDK获取设备能力值 */
type YoukuOttAlicbFacadeserviceGetdataAPIRequest struct {
	model.Params
	// 能力维度
	_serviceList []string
	// 设备唯一标识
	_uuid string
	// 属性MAP JSON串
	_propertyMapJson string
	// 扩展属性
	_extraInfoMap string
}

// New
