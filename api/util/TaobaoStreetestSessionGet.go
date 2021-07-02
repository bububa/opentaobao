package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoStreetestSessionGet 根据获取压测用户的sessionKey
// taobao.streetest.session.get
//
// 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
func TaobaoStreetestSessionGet(clt *core.SDKClient, req *util.TaobaoStreetestSessionGetAPIRequest, session string) (*util.TaobaoStreetestSessionGetAPIResponse, error) {
	var resp util.TaobaoStreetestSessionGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
