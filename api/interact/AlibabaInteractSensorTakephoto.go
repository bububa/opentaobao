package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorTakephoto takePhoto
// alibaba.interact.sensor.takephoto
//
// 客户端takePhoto
func AlibabaInteractSensorTakephoto(clt *core.SDKClient, req *interact.AlibabaInteractSensorTakephotoAPIRequest, resp *interact.AlibabaInteractSensorTakephotoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
