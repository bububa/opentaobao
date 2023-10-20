package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseAdminThemeCreateAPIRequest 创建云主题 API请求
// alibaba.alihouse.admin.theme.create
//
// 创建云主题
type AlibabaAlihouseAdminThemeCreateAPIRequest struct {
	model.Params
	// 请求云主题参数
	_etcThemeDto *EtcThemeDto
}

// NewAlibabaAlihouseAdminThemeCreateRequest 初始化AlibabaAlihouseAdminThemeCreateAPIRequest对象
func NewAlibabaAlihouseAdminThemeCreateRequest() *AlibabaAlihouseAdminThemeCreateAPIRequest {
	return &AlibabaAlihouseAdminThemeCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseAdminThemeCreateAPIRequest) Reset() {
	r._etcThemeDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseAdminThemeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.admin.theme.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseAdminThemeCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseAdminThemeCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEtcThemeDto is EtcThemeDto Setter
// 请求云主题参数
func (r *AlibabaAlihouseAdminThemeCreateAPIRequest) SetEtcThemeDto(_etcThemeDto *EtcThemeDto) error {
	r._etcThemeDto = _etcThemeDto
	r.Set("etc_theme_dto", _etcThemeDto)
	return nil
}

// GetEtcThemeDto EtcThemeDto Getter
func (r AlibabaAlihouseAdminThemeCreateAPIRequest) GetEtcThemeDto() *EtcThemeDto {
	return r._etcThemeDto
}

var poolAlibabaAlihouseAdminThemeCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseAdminThemeCreateRequest()
	},
}

// GetAlibabaAlihouseAdminThemeCreateRequest 从 sync.Pool 获取 AlibabaAlihouseAdminThemeCreateAPIRequest
func GetAlibabaAlihouseAdminThemeCreateAPIRequest() *AlibabaAlihouseAdminThemeCreateAPIRequest {
	return poolAlibabaAlihouseAdminThemeCreateAPIRequest.Get().(*AlibabaAlihouseAdminThemeCreateAPIRequest)
}

// ReleaseAlibabaAlihouseAdminThemeCreateAPIRequest 将 AlibabaAlihouseAdminThemeCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseAdminThemeCreateAPIRequest(v *AlibabaAlihouseAdminThemeCreateAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseAdminThemeCreateAPIRequest.Put(v)
}
