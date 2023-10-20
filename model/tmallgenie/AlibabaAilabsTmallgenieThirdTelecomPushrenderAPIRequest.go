package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest 电信-推送拉起设备应用 API请求
// alibaba.ailabs.tmallgenie.third.telecom.pushrender
//
// 电信-推送拉起设备应用
type AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest struct {
	model.Params
	// ctei
	_ctei string
	// 拉起端上应用/服务的url
	_url string
	// 额外传递的参数
	_params string
	// 应用拉起方式
	_model string
}

// NewAlibabaailabstmallgeniethirdtelecompushrenderRequest 初始化AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest对象
func NewAlibabaailabstmallgeniethirdtelecompushrenderRequest() *AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest {
	return &AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.third.telecom.pushrender"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCtei is Ctei Setter
// ctei
func (r *AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) SetCtei(_ctei string) error {
	r._ctei = _ctei
	r.Set("ctei", _ctei)
	return nil
}

// GetCtei Ctei Getter
func (r AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) GetCtei() string {
	return r._ctei
}

// SetUrl is Url Setter
// 拉起端上应用/服务的url
func (r *AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) GetUrl() string {
	return r._url
}

// SetParams is Params Setter
// 额外传递的参数
func (r *AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) GetParams() string {
	return r._params
}

// SetModel is Model Setter
// 应用拉起方式
func (r *AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) SetModel(_model string) error {
	r._model = _model
	r.Set("model", _model)
	return nil
}

// GetModel Model Getter
func (r AlibabaailabstmallgeniethirdtelecompushrenderAPIRequest) GetModel() string {
	return r._model
}
