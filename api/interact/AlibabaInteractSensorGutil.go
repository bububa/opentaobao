package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGutil canvas工具包
// alibaba.interact.sensor.gutil
//
// canvas工具包
func AlibabaInteractSensorGutil(clt *core.SDKClient, req *interact.AlibabaInteractSensorGutilAPIRequest, resp *interact.AlibabaInteractSensorGutilAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
