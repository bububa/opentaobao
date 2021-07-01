package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceGetdeviceforqueryAPIRequest
下发设备的分页接口(无需AOP控制) API请求
alibaba.campus.device.getdeviceforquery

下发设备的分页接口(发布在TOP上，connect调用，无需AOP控制) */
type AlibabaCampusDeviceGetdeviceforqueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_query *DeviceApiQuery
	// 平台统一参数
	_workBenchContext *WorkBenchContext
}

// New
