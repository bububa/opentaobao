package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorWangwangAPIRequest
手淘拉起旺旺接口 API请求
alibaba.interact.sensor.wangwang

手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。手淘开放旺旺拉起功能给ISV。 */
type AlibabaInteractSensorWangwangAPIRequest struct {
	model.Params
}

// New
