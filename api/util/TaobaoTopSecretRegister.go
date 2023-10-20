package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotopsecretregister 注册加密账号
// taobao.top.secret.register
//
// 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
func Taobaotopsecretregister(clt *core.SDKClient, req *util.TaobaotopsecretregisterAPIRequest, session string) (*util.TaobaotopsecretregisterAPIResponse, error) {
	var resp util.TaobaotopsecretregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
