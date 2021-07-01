package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthUicUserinfoHealthidGetAPIRequest
获取健康id API请求
alibaba.alihealth.uic.userinfo.healthid.get

根据支付宝用户ID获取用户健康ID */
type AlibabaAlihealthUicUserinfoHealthidGetAPIRequest struct {
	model.Params
	// 支付宝用户ID
	_alipayUserId string
}

// New
