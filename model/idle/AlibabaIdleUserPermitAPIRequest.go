package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleUserPermitAPIRequest
用户appkey授权 API请求
alibaba.idle.user.permit

用于记录登录用户与服务商的绑定关系，用于业务数据分发和授权校验 */
type AlibabaIdleUserPermitAPIRequest struct {
	model.Params
	// 授权请求
	_paramUserGrantRequest *UserGrantRequest
}

// New
