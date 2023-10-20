package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopauthlogout 登出
// taobao.ailab.aicloud.top.auth.logout
//
// 登出
func Taobaoailabaicloudtopauthlogout(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopauthlogoutAPIRequest, session string) (*iot.TaobaoailabaicloudtopauthlogoutAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopauthlogoutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
