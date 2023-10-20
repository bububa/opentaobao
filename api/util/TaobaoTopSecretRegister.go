package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTopSecretRegister 注册加密账号
// taobao.top.secret.register
//
// 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
func TaobaoTopSecretRegister(clt *core.SDKClient, req *util.TaobaoTopSecretRegisterAPIRequest, resp *util.TaobaoTopSecretRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
