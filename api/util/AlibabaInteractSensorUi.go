package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// AlibabaInteractSensorUi 基本ui操作
// alibaba.interact.sensor.ui
//
// Weex 基本UI操作
func AlibabaInteractSensorUi(clt *core.SDKClient, req *util.AlibabaInteractSensorUiAPIRequest, resp *util.AlibabaInteractSensorUiAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
