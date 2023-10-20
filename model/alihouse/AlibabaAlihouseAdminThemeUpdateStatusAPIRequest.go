package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseAdminThemeUpdateStatusAPIRequest 云主题上下架+删除 API请求
// alibaba.alihouse.admin.theme.update.status
//
// 云主题上下架+删除
type AlibabaAlihouseAdminThemeUpdateStatusAPIRequest struct {
	model.Params
	// 请求云主题参数
	_etcThemeDto *EtcThemeDto
}

// NewAlibabaAlihouseAdminThemeUpdateStatusRequest 初始化AlibabaAlihouseAdminThemeUpdateStatusAPIRequest对象
func NewAlibabaAlihouseAdminThemeUpdateStatusRequest() *AlibabaAlihouseAdminThemeUpdateStatusAPIRequest {
	return &AlibabaAlihouseAdminThemeUpdateStatusAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseAdminThemeUpdateStatusAPIRequest) Reset() {
	r._etcThemeDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseAdminThemeUpdateStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.admin.theme.update.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseAdminThemeUpdateStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseAdminThemeUpdateStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEtcThemeDto is EtcThemeDto Setter
// 请求云主题参数
func (r *AlibabaAlihouseAdminThemeUpdateStatusAPIRequest) SetEtcThemeDto(_etcThemeDto *EtcThemeDto) error {
	r._etcThemeDto = _etcThemeDto
	r.Set("etc_theme_dto", _etcThemeDto)
	return nil
}

// GetEtcThemeDto EtcThemeDto Getter
func (r AlibabaAlihouseAdminThemeUpdateStatusAPIRequest) GetEtcThemeDto() *EtcThemeDto {
	return r._etcThemeDto
}

var poolAlibabaAlihouseAdminThemeUpdateStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseAdminThemeUpdateStatusRequest()
	},
}

// GetAlibabaAlihouseAdminThemeUpdateStatusRequest 从 sync.Pool 获取 AlibabaAlihouseAdminThemeUpdateStatusAPIRequest
func GetAlibabaAlihouseAdminThemeUpdateStatusAPIRequest() *AlibabaAlihouseAdminThemeUpdateStatusAPIRequest {
	return poolAlibabaAlihouseAdminThemeUpdateStatusAPIRequest.Get().(*AlibabaAlihouseAdminThemeUpdateStatusAPIRequest)
}

// ReleaseAlibabaAlihouseAdminThemeUpdateStatusAPIRequest 将 AlibabaAlihouseAdminThemeUpdateStatusAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseAdminThemeUpdateStatusAPIRequest(v *AlibabaAlihouseAdminThemeUpdateStatusAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseAdminThemeUpdateStatusAPIRequest.Put(v)
}
