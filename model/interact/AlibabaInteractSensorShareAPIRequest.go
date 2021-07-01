package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorShareAPIRequest
分享 API请求
alibaba.interact.sensor.share

客户端分享 */
type AlibabaInteractSensorShareAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorShareRequest 初始化AlibabaInteractSensorShareAPIRequest对象
func NewAlibabaInteractSensorShareRequest() *AlibabaInteractSensorShareAPIRequest {
	return &AlibabaInteractSensorShareAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorShareAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.share"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorShareAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
