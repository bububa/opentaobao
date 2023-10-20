package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorgutil canvas工具包
// alibaba.interact.sensor.gutil
//
// canvas工具包
func Alibabainteractsensorgutil(clt *core.SDKClient, req *interact.AlibabainteractsensorgutilAPIRequest, session string) (*interact.AlibabainteractsensorgutilAPIResponse, error) {
	var resp interact.AlibabainteractsensorgutilAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
