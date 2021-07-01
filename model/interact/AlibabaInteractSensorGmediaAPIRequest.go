package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorGmediaAPIRequest
gmedia API请求
alibaba.interact.sensor.gmedia

媒体功能 */
type AlibabaInteractSensorGmediaAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGmediaRequest 初始化AlibabaInteractSensorGmediaAPIRequest对象
func NewAlibabaInteractSensorGmediaRequest() *AlibabaInteractSensorGmediaAPIRequest {
	return &AlibabaInteractSensorGmediaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGmediaAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gmedia"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGmediaAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
