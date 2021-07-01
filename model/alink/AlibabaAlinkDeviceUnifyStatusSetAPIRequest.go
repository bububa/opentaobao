package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlinkDeviceUnifyStatusSetAPIRequest
设置设备标准属性状态 API请求
alibaba.alink.device.unify.status.set

操作用户绑定的设备 */
type AlibabaAlinkDeviceUnifyStatusSetAPIRequest struct {
	model.Params
	// uuid
	_uuid string
	// 设备的设置参数数据
	_instructions string
}

// New
