package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorWangwang 手淘拉起旺旺接口
// alibaba.interact.sensor.wangwang
//
// 手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。手淘开放旺旺拉起功能给ISV。
func AlibabaInteractSensorWangwang(clt *core.SDKClient, req *interact.AlibabaInteractSensorWangwangAPIRequest, session string) (*interact.AlibabaInteractSensorWangwangAPIResponse, error) {
	var resp interact.AlibabaInteractSensorWangwangAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
