package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaloginauthsignget 获取登陆权限签名
// taobao.simba.login.authsign.get
//
// 获取登陆权限签名
func Taobaosimbaloginauthsignget(clt *core.SDKClient, req *simba.TaobaosimbaloginauthsigngetAPIRequest, session string) (*simba.TaobaosimbaloginauthsigngetAPIResponse, error) {
	var resp simba.TaobaosimbaloginauthsigngetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
