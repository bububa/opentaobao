package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVasServiceValidateAPIRequest
增值服务订购服务验证 API请求
taobao.vas.service.validate

增值服务订购服务验证 */
type TaobaoVasServiceValidateAPIRequest struct {
	model.Params
	// 服务编码
	_servCode string
	// 用户昵称
	_nick string
}

// New
