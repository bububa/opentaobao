package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorlogin 获取登陆页面
// alibaba.interact.sensor.login
//
// 获取登陆页面
func Alibabainteractsensorlogin(clt *core.SDKClient, req *interact.AlibabainteractsensorloginAPIRequest, session string) (*interact.AlibabainteractsensorloginAPIResponse, error) {
	var resp interact.AlibabainteractsensorloginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
