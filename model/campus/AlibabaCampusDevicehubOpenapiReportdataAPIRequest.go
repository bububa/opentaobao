package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdevicehubopenapireportdataAPIRequest 设备数据上报 API请求
// alibaba.campus.devicehub.openapi.reportdata
//
// 设备数据上报
type AlibabacampusdevicehubopenapireportdataAPIRequest struct {
	model.Params
	// 设备数据上报信息
	_deviceEventData *DeviceReportEventDto
}

// NewAlibabacampusdevicehubopenapireportdataRequest 初始化AlibabacampusdevicehubopenapireportdataAPIRequest对象
func NewAlibabacampusdevicehubopenapireportdataRequest() *AlibabacampusdevicehubopenapireportdataAPIRequest {
	return &AlibabacampusdevicehubopenapireportdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdevicehubopenapireportdataAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.devicehub.openapi.reportdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdevicehubopenapireportdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdevicehubopenapireportdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceEventData is DeviceEventData Setter
// 设备数据上报信息
func (r *AlibabacampusdevicehubopenapireportdataAPIRequest) SetDeviceEventData(_deviceEventData *DeviceReportEventDto) error {
	r._deviceEventData = _deviceEventData
	r.Set("device_event_data", _deviceEventData)
	return nil
}

// GetDeviceEventData DeviceEventData Getter
func (r AlibabacampusdevicehubopenapireportdataAPIRequest) GetDeviceEventData() *DeviceReportEventDto {
	return r._deviceEventData
}
