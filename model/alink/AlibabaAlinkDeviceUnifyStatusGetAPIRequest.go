package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceUnifyStatusGetAPIRequest 查询设备标准属性最新状态 API请求
// alibaba.alink.device.unify.status.get
//
// 查询设备最新标准属性状态
type AlibabaAlinkDeviceUnifyStatusGetAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
}

// NewAlibabaAlinkDeviceUnifyStatusGetRequest 初始化AlibabaAlinkDeviceUnifyStatusGetAPIRequest对象
func NewAlibabaAlinkDeviceUnifyStatusGetRequest() *AlibabaAlinkDeviceUnifyStatusGetAPIRequest {
	return &AlibabaAlinkDeviceUnifyStatusGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceUnifyStatusGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.unify.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceUnifyStatusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlinkDeviceUnifyStatusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceUnifyStatusGetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAlinkDeviceUnifyStatusGetAPIRequest) GetUuid() string {
	return r._uuid
}
