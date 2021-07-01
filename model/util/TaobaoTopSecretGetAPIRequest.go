package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopSecretGetAPIRequest
获取TOP通道解密秘钥 API请求
taobao.top.secret.get

top sdk通过api获取对应解密秘钥 */
type TaobaoTopSecretGetAPIRequest struct {
	model.Params
	// 秘钥版本号
	_secretVersion int64
	// 伪随机数
	_randomNum string
	// 自定义用户id
	_customerUserId int64
}

// New
