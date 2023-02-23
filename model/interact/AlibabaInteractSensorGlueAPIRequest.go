package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGlueAPIRequest 视频播放 API请求
// alibaba.interact.sensor.glue
//
// 视频播放
type AlibabaInteractSensorGlueAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGlueRequest 初始化AlibabaInteractSensorGlueAPIRequest对象
func NewAlibabaInteractSensorGlueRequest() *AlibabaInteractSensorGlueAPIRequest {
	return &AlibabaInteractSensorGlueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGlueAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.glue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGlueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorGlueAPIRequest) GetRawParams() model.Params {
	return r.Params
}
