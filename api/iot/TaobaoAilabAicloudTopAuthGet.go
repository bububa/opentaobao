package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopauthget 登陆
// taobao.ailab.aicloud.top.auth.get
//
// 登陆
func Taobaoailabaicloudtopauthget(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopauthgetAPIRequest, session string) (*iot.TaobaoailabaicloudtopauthgetAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopauthgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
