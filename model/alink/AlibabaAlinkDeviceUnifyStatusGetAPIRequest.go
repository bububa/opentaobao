package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlinkDeviceUnifyStatusGetAPIRequest
查询设备标准属性最新状态 API请求
alibaba.alink.device.unify.status.get

查询设备最新标准属性状态 */
type AlibabaAlinkDeviceUnifyStatusGetAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// New
