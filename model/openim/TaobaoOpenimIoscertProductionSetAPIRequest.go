package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimIoscertProductionSetAPIRequest
设置ios证书 API请求
taobao.openim.ioscert.production.set

设置ios证书 */
type TaobaoOpenimIoscertProductionSetAPIRequest struct {
	model.Params
	// 证书密码
	_password string
	// 证书文件内容,base64编码
	_cert string
}

// New
