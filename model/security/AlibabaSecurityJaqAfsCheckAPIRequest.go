package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqAfsCheckAPIRequest
反欺诈二次验证接口 API请求
alibaba.security.jaq.afs.check

反欺诈二次验证接口 */
type AlibabaSecurityJaqAfsCheckAPIRequest struct {
	model.Params
	// 上报平台枚举值 1标识Android端 2标识iOS端 3标识PC端及其他
	_platform int64
	// token，来自客户端上报
	_token string
	// 会话ID，来自客户端上报
	_sessionId string
	// 签名串，来自客户端上报
	_sig string
	// 用户接入的时候获取的风控key
	_afsKey string
}

// New
