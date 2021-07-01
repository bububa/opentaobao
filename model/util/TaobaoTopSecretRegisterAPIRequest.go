package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopSecretRegisterAPIRequest
注册加密账号 API请求
taobao.top.secret.register

提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密 */
type TaobaoTopSecretRegisterAPIRequest struct {
	model.Params
	// 用户id，保证唯一
	_userId int64
}

// New
