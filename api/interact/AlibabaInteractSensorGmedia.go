package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorgmedia gmedia
// alibaba.interact.sensor.gmedia
//
// 媒体功能
func Alibabainteractsensorgmedia(clt *core.SDKClient, req *interact.AlibabainteractsensorgmediaAPIRequest, session string) (*interact.AlibabainteractsensorgmediaAPIResponse, error) {
	var resp interact.AlibabainteractsensorgmediaAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
