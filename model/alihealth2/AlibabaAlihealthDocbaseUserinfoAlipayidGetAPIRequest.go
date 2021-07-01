package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest
根据健康ID获取支付宝ID API请求
alibaba.alihealth.docbase.userinfo.alipayid.get

根据健康ID获取支付宝ID */
type AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest struct {
	model.Params
	// 阿里健康ID
	_alihealthUserId string
	// 渠道alipay taobao uc gaode
	_appChannel string
}

// New
