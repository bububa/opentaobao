package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingServicepackAPIRequest
获取海王用户权限包 API请求
alibaba.seaking.servicepack

获取海王用户权限包 */
type AlibabaSeakingServicepackAPIRequest struct {
	model.Params
	// 验证类型
	_identifyType string
	// 验证类型下的唯一id
	_identifier string
}

// New
