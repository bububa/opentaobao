package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotopopenidconvert 混淆nick转openid
// taobao.top.openid.convert
//
// 混淆nick转openid，生成混淆nick必须与当前请求的isv匹配
func Taobaotopopenidconvert(clt *core.SDKClient, req *util.TaobaotopopenidconvertAPIRequest, session string) (*util.TaobaotopopenidconvertAPIResponse, error) {
	var resp util.TaobaotopopenidconvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
