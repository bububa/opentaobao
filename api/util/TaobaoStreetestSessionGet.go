package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaostreetestsessionget 根据获取压测用户的sessionKey
// taobao.streetest.session.get
//
// 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
func Taobaostreetestsessionget(clt *core.SDKClient, req *util.TaobaostreetestsessiongetAPIRequest, session string) (*util.TaobaostreetestsessiongetAPIResponse, error) {
	var resp util.TaobaostreetestsessiongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
