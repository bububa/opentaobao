package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest 是否支持云回看新SDK API请求
// alibaba.ailabs.tmallgenie.sdk.device.issupportsdk
//
// 是否支持云回看新SDK
type AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest struct {
	model.Params
	// 天翼账号
	_tyAccount string
	// 设备的ctei串码
	_ctei string
}

// NewAlibabaailabstmallgeniesdkdeviceissupportsdkRequest 初始化AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest对象
func NewAlibabaailabstmallgeniesdkdeviceissupportsdkRequest() *AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest {
	return &AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.sdk.device.issupportsdk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTyAccount is TyAccount Setter
// 天翼账号
func (r *AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest) SetTyAccount(_tyAccount string) error {
	r._tyAccount = _tyAccount
	r.Set("ty_account", _tyAccount)
	return nil
}

// GetTyAccount TyAccount Getter
func (r AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest) GetTyAccount() string {
	return r._tyAccount
}

// SetCtei is Ctei Setter
// 设备的ctei串码
func (r *AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest) SetCtei(_ctei string) error {
	r._ctei = _ctei
	r.Set("ctei", _ctei)
	return nil
}

// GetCtei Ctei Getter
func (r AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest) GetCtei() string {
	return r._ctei
}
