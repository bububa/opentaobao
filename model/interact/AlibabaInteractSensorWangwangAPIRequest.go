package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorWangwangAPIRequest 手淘拉起旺旺接口 API请求
// alibaba.interact.sensor.wangwang
//
// 手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。手淘开放旺旺拉起功能给ISV。
type AlibabaInteractSensorWangwangAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorWangwangRequest 初始化AlibabaInteractSensorWangwangAPIRequest对象
func NewAlibabaInteractSensorWangwangRequest() *AlibabaInteractSensorWangwangAPIRequest {
	return &AlibabaInteractSensorWangwangAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorWangwangAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.wangwang"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorWangwangAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
