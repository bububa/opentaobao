package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceUnifyStatusSetAPIRequest 设置设备标准属性状态 API请求
// alibaba.alink.device.unify.status.set
//
// 操作用户绑定的设备
type AlibabaAlinkDeviceUnifyStatusSetAPIRequest struct {
	model.Params
	// uuid
	_uuid string
	// 设备的设置参数数据
	_instructions string
}

// NewAlibabaAlinkDeviceUnifyStatusSetRequest 初始化AlibabaAlinkDeviceUnifyStatusSetAPIRequest对象
func NewAlibabaAlinkDeviceUnifyStatusSetRequest() *AlibabaAlinkDeviceUnifyStatusSetAPIRequest {
	return &AlibabaAlinkDeviceUnifyStatusSetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceUnifyStatusSetAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.unify.status.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceUnifyStatusSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlinkDeviceUnifyStatusSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// uuid
func (r *AlibabaAlinkDeviceUnifyStatusSetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAlinkDeviceUnifyStatusSetAPIRequest) GetUuid() string {
	return r._uuid
}

// SetInstructions is Instructions Setter
// 设备的设置参数数据
func (r *AlibabaAlinkDeviceUnifyStatusSetAPIRequest) SetInstructions(_instructions string) error {
	r._instructions = _instructions
	r.Set("instructions", _instructions)
	return nil
}

// GetInstructions Instructions Getter
func (r AlibabaAlinkDeviceUnifyStatusSetAPIRequest) GetInstructions() string {
	return r._instructions
}
