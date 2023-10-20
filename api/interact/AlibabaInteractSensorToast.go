package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensortoast toast
// alibaba.interact.sensor.toast
//
// toast提示
func Alibabainteractsensortoast(clt *core.SDKClient, req *interact.AlibabainteractsensortoastAPIRequest, session string) (*interact.AlibabainteractsensortoastAPIResponse, error) {
	var resp interact.AlibabainteractsensortoastAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
