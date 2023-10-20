package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotmcauthget TMC授权token
// taobao.tmc.auth.get
//
// TMC连接授权Token
func Taobaotmcauthget(clt *core.SDKClient, req *util.TaobaotmcauthgetAPIRequest, session string) (*util.TaobaotmcauthgetAPIResponse, error) {
	var resp util.TaobaotmcauthgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
