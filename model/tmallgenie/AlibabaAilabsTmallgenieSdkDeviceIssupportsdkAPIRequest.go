package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest 是否支持云回看新SDK API请求
// alibaba.ailabs.tmallgenie.sdk.device.issupportsdk
//
// 是否支持云回看新SDK
type AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest struct {
	model.Params
	// 天翼账号
	_tyAccount string
	// 设备的ctei串码
	_ctei string
}

// NewAlibabaAilabsTmallgenieSdkDeviceIssupportsdkRequest 初始化AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest对象
func NewAlibabaAilabsTmallgenieSdkDeviceIssupportsdkRequest() *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest {
	return &AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest) Reset() {
	r._tyAccount = ""
	r._ctei = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.sdk.device.issupportsdk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTyAccount is TyAccount Setter
// 天翼账号
func (r *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest) SetTyAccount(_tyAccount string) error {
	r._tyAccount = _tyAccount
	r.Set("ty_account", _tyAccount)
	return nil
}

// GetTyAccount TyAccount Getter
func (r AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest) GetTyAccount() string {
	return r._tyAccount
}

// SetCtei is Ctei Setter
// 设备的ctei串码
func (r *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest) SetCtei(_ctei string) error {
	r._ctei = _ctei
	r.Set("ctei", _ctei)
	return nil
}

// GetCtei Ctei Getter
func (r AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest) GetCtei() string {
	return r._ctei
}

var poolAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsTmallgenieSdkDeviceIssupportsdkRequest()
	},
}

// GetAlibabaAilabsTmallgenieSdkDeviceIssupportsdkRequest 从 sync.Pool 获取 AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest
func GetAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest() *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest {
	return poolAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest.Get().(*AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest)
}

// ReleaseAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest 将 AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest(v *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest) {
	v.Reset()
	poolAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest.Put(v)
}
