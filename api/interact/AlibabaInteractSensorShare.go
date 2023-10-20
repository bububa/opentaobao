package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorshare 分享
// alibaba.interact.sensor.share
//
// 客户端分享
func Alibabainteractsensorshare(clt *core.SDKClient, req *interact.AlibabainteractsensorshareAPIRequest, session string) (*interact.AlibabainteractsensorshareAPIResponse, error) {
	var resp interact.AlibabainteractsensorshareAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
