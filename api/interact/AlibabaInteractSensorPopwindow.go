package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorpopwindow popwindow
// alibaba.interact.sensor.popwindow
//
// popwindow
func Alibabainteractsensorpopwindow(clt *core.SDKClient, req *interact.AlibabainteractsensorpopwindowAPIRequest, session string) (*interact.AlibabainteractsensorpopwindowAPIResponse, error) {
	var resp interact.AlibabainteractsensorpopwindowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
