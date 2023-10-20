package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicegetcodeAPIRequest 获取authcode API请求
// alibaba.ailabs.tmallgenie.auth.device.getcode
//
// 获取绑定的authcode
type AlibabaailabstmallgenieauthdevicegetcodeAPIRequest struct {
	model.Params
}

// NewAlibabaailabstmallgenieauthdevicegetcodeRequest 初始化AlibabaailabstmallgenieauthdevicegetcodeAPIRequest对象
func NewAlibabaailabstmallgenieauthdevicegetcodeRequest() *AlibabaailabstmallgenieauthdevicegetcodeAPIRequest {
	return &AlibabaailabstmallgenieauthdevicegetcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdevicegetcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.getcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdevicegetcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdevicegetcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}
