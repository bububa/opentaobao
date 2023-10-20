package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest 根据ctei查询状态 API请求
// alibaba.ailabs.tmallgenie.auth.device.status.getbyctei
//
// 提供给电信查询设备在线状态值
type AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest struct {
	model.Params
	// ctei
	_ctei string
}

// NewAlibabaailabstmallgenieauthdevicestatusgetbycteiRequest 初始化AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest对象
func NewAlibabaailabstmallgenieauthdevicestatusgetbycteiRequest() *AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest {
	return &AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.status.getbyctei"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCtei is Ctei Setter
// ctei
func (r *AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest) SetCtei(_ctei string) error {
	r._ctei = _ctei
	r.Set("ctei", _ctei)
	return nil
}

// GetCtei Ctei Getter
func (r AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest) GetCtei() string {
	return r._ctei
}
