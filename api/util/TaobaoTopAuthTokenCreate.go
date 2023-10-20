package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotopauthtokencreate 获取Access Token
// taobao.top.auth.token.create
//
// 用户通过code换获取access_token，https only
func Taobaotopauthtokencreate(clt *core.SDKClient, req *util.TaobaotopauthtokencreateAPIRequest, session string) (*util.TaobaotopauthtokencreateAPIResponse, error) {
	var resp util.TaobaotopauthtokencreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
