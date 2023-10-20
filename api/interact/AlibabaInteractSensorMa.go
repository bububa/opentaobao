package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorma 码相关API
// alibaba.interact.sensor.ma
//
// 码相关API
func Alibabainteractsensorma(clt *core.SDKClient, req *interact.AlibabainteractsensormaAPIRequest, session string) (*interact.AlibabainteractsensormaAPIResponse, error) {
	var resp interact.AlibabainteractsensormaAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
