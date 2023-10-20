package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorblow 吹气
// alibaba.interact.sensor.blow
//
// 客户端吹气
func Alibabainteractsensorblow(clt *core.SDKClient, req *interact.AlibabainteractsensorblowAPIRequest, session string) (*interact.AlibabainteractsensorblowAPIResponse, error) {
	var resp interact.AlibabainteractsensorblowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
