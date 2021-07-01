package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunAliosPayTokenGetAPIRequest
获取支付token API请求
aliyun.alios.pay.token.get

商户用来获取支付的授权token */
type AliyunAliosPayTokenGetAPIRequest struct {
	model.Params
	// 请求参数
	_getTokenRequest *GetTokenRequest
}

// New
