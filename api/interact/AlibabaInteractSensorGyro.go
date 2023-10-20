package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorgyro 陀螺仪
// alibaba.interact.sensor.gyro
//
// 客户端陀螺仪
func Alibabainteractsensorgyro(clt *core.SDKClient, req *interact.AlibabainteractsensorgyroAPIRequest, session string) (*interact.AlibabainteractsensorgyroAPIResponse, error) {
	var resp interact.AlibabainteractsensorgyroAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
