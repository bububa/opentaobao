package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGyroAPIRequest 陀螺仪 API请求
// alibaba.interact.sensor.gyro
//
// 客户端陀螺仪
type AlibabaInteractSensorGyroAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGyroRequest 初始化AlibabaInteractSensorGyroAPIRequest对象
func NewAlibabaInteractSensorGyroRequest() *AlibabaInteractSensorGyroAPIRequest {
	return &AlibabaInteractSensorGyroAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGyroAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gyro"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGyroAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
