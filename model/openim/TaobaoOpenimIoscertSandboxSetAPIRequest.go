package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimIoscertSandboxSetAPIRequest
设置开发环境证书 API请求
taobao.openim.ioscert.sandbox.set

设置开发环境证书 */
type TaobaoOpenimIoscertSandboxSetAPIRequest struct {
	model.Params
	// 证书内容,base64编码
	_cert string
	// 系统自动生成
	_password string
}

// New
