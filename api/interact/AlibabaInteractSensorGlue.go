package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

/* AlibabaInteractSensorGlue
视频播放
alibaba.interact.sensor.glue

视频播放 */
func AlibabaInteractSensorGlue(clt *core.SDKClient, req *interact.AlibabaInteractSensorGlueAPIRequest, session string) (*interact.AlibabaInteractSensorGlueAPIResponse, error) {
	var resp interact.AlibabaInteractSensorGlueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
