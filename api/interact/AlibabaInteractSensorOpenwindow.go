package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensoropenwindow 客户端打开新页面
// alibaba.interact.sensor.openwindow
//
// 客户端打开新页面
func Alibabainteractsensoropenwindow(clt *core.SDKClient, req *interact.AlibabainteractsensoropenwindowAPIRequest, session string) (*interact.AlibabainteractsensoropenwindowAPIResponse, error) {
	var resp interact.AlibabainteractsensoropenwindowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
