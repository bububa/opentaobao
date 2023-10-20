package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTmsCutConfirmAPIRequest 配拦截失败CP确认结果并回告 API请求
// alibaba.alihealth.tms.cut.confirm
//
// 配拦截失败CP确认结果并回告
type AlibabaAlihealthTmsCutConfirmAPIRequest struct {
	model.Params
	// 参数
	_tmsCutResultConfirmRequest *TmsCutResultConfirmRequest
}

// NewAlibabaAlihealthTmsCutConfirmRequest 初始化AlibabaAlihealthTmsCutConfirmAPIRequest对象
func NewAlibabaAlihealthTmsCutConfirmRequest() *AlibabaAlihealthTmsCutConfirmAPIRequest {
	return &AlibabaAlihealthTmsCutConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTmsCutConfirmAPIRequest) Reset() {
	r._tmsCutResultConfirmRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTmsCutConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tms.cut.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTmsCutConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTmsCutConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsCutResultConfirmRequest is TmsCutResultConfirmRequest Setter
// 参数
func (r *AlibabaAlihealthTmsCutConfirmAPIRequest) SetTmsCutResultConfirmRequest(_tmsCutResultConfirmRequest *TmsCutResultConfirmRequest) error {
	r._tmsCutResultConfirmRequest = _tmsCutResultConfirmRequest
	r.Set("tms_cut_result_confirm_request", _tmsCutResultConfirmRequest)
	return nil
}

// GetTmsCutResultConfirmRequest TmsCutResultConfirmRequest Getter
func (r AlibabaAlihealthTmsCutConfirmAPIRequest) GetTmsCutResultConfirmRequest() *TmsCutResultConfirmRequest {
	return r._tmsCutResultConfirmRequest
}

var poolAlibabaAlihealthTmsCutConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTmsCutConfirmRequest()
	},
}

// GetAlibabaAlihealthTmsCutConfirmRequest 从 sync.Pool 获取 AlibabaAlihealthTmsCutConfirmAPIRequest
func GetAlibabaAlihealthTmsCutConfirmAPIRequest() *AlibabaAlihealthTmsCutConfirmAPIRequest {
	return poolAlibabaAlihealthTmsCutConfirmAPIRequest.Get().(*AlibabaAlihealthTmsCutConfirmAPIRequest)
}

// ReleaseAlibabaAlihealthTmsCutConfirmAPIRequest 将 AlibabaAlihealthTmsCutConfirmAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTmsCutConfirmAPIRequest(v *AlibabaAlihealthTmsCutConfirmAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTmsCutConfirmAPIRequest.Put(v)
}
