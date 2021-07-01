package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

/* AlibabaInteractSensorMa
码相关API
alibaba.interact.sensor.ma

码相关API */
func AlibabaInteractSensorMa(clt *core.SDKClient, req *interact.AlibabaInteractSensorMaAPIRequest, session string) (*interact.AlibabaInteractSensorMaAPIResponse, error) {
	var resp interact.AlibabaInteractSensorMaAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
