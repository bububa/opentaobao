package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorBlowAPIRequest
吹气 API请求
alibaba.interact.sensor.blow

客户端吹气 */
type AlibabaInteractSensorBlowAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorBlowRequest 初始化AlibabaInteractSensorBlowAPIRequest对象
func NewAlibabaInteractSensorBlowRequest() *AlibabaInteractSensorBlowAPIRequest {
	return &AlibabaInteractSensorBlowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorBlowAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.blow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorBlowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
