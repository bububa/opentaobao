package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseAdminThemeUpdateAPIRequest 云主题更新 API请求
// alibaba.alihouse.admin.theme.update
//
// 云主题更新
type AlibabaAlihouseAdminThemeUpdateAPIRequest struct {
	model.Params
	// 请求云主题参数
	_etcThemeDto *EtcThemeDto
}

// NewAlibabaAlihouseAdminThemeUpdateRequest 初始化AlibabaAlihouseAdminThemeUpdateAPIRequest对象
func NewAlibabaAlihouseAdminThemeUpdateRequest() *AlibabaAlihouseAdminThemeUpdateAPIRequest {
	return &AlibabaAlihouseAdminThemeUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseAdminThemeUpdateAPIRequest) Reset() {
	r._etcThemeDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseAdminThemeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.admin.theme.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseAdminThemeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseAdminThemeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEtcThemeDto is EtcThemeDto Setter
// 请求云主题参数
func (r *AlibabaAlihouseAdminThemeUpdateAPIRequest) SetEtcThemeDto(_etcThemeDto *EtcThemeDto) error {
	r._etcThemeDto = _etcThemeDto
	r.Set("etc_theme_dto", _etcThemeDto)
	return nil
}

// GetEtcThemeDto EtcThemeDto Getter
func (r AlibabaAlihouseAdminThemeUpdateAPIRequest) GetEtcThemeDto() *EtcThemeDto {
	return r._etcThemeDto
}

var poolAlibabaAlihouseAdminThemeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseAdminThemeUpdateRequest()
	},
}

// GetAlibabaAlihouseAdminThemeUpdateRequest 从 sync.Pool 获取 AlibabaAlihouseAdminThemeUpdateAPIRequest
func GetAlibabaAlihouseAdminThemeUpdateAPIRequest() *AlibabaAlihouseAdminThemeUpdateAPIRequest {
	return poolAlibabaAlihouseAdminThemeUpdateAPIRequest.Get().(*AlibabaAlihouseAdminThemeUpdateAPIRequest)
}

// ReleaseAlibabaAlihouseAdminThemeUpdateAPIRequest 将 AlibabaAlihouseAdminThemeUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseAdminThemeUpdateAPIRequest(v *AlibabaAlihouseAdminThemeUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseAdminThemeUpdateAPIRequest.Put(v)
}
