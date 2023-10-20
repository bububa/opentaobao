package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorgyroAPIRequest 陀螺仪 API请求
// alibaba.interact.sensor.gyro
//
// 客户端陀螺仪
type AlibabainteractsensorgyroAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorgyroRequest 初始化AlibabainteractsensorgyroAPIRequest对象
func NewAlibabainteractsensorgyroRequest() *AlibabainteractsensorgyroAPIRequest {
	return &AlibabainteractsensorgyroAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorgyroAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gyro"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorgyroAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorgyroAPIRequest) GetRawParams() model.Params {
	return r.Params
}
