package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseAdminThemeHotUpdateAPIRequest 云主题热更新数据集 API请求
// alibaba.alihouse.admin.theme.hot.update
//
// 云主题更新
type AlibabaAlihouseAdminThemeHotUpdateAPIRequest struct {
	model.Params
	// 请求云主题参数
	_etcThemeDto *EtcThemeDto
}

// NewAlibabaAlihouseAdminThemeHotUpdateRequest 初始化AlibabaAlihouseAdminThemeHotUpdateAPIRequest对象
func NewAlibabaAlihouseAdminThemeHotUpdateRequest() *AlibabaAlihouseAdminThemeHotUpdateAPIRequest {
	return &AlibabaAlihouseAdminThemeHotUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseAdminThemeHotUpdateAPIRequest) Reset() {
	r._etcThemeDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseAdminThemeHotUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.admin.theme.hot.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseAdminThemeHotUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseAdminThemeHotUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEtcThemeDto is EtcThemeDto Setter
// 请求云主题参数
func (r *AlibabaAlihouseAdminThemeHotUpdateAPIRequest) SetEtcThemeDto(_etcThemeDto *EtcThemeDto) error {
	r._etcThemeDto = _etcThemeDto
	r.Set("etc_theme_dto", _etcThemeDto)
	return nil
}

// GetEtcThemeDto EtcThemeDto Getter
func (r AlibabaAlihouseAdminThemeHotUpdateAPIRequest) GetEtcThemeDto() *EtcThemeDto {
	return r._etcThemeDto
}

var poolAlibabaAlihouseAdminThemeHotUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseAdminThemeHotUpdateRequest()
	},
}

// GetAlibabaAlihouseAdminThemeHotUpdateRequest 从 sync.Pool 获取 AlibabaAlihouseAdminThemeHotUpdateAPIRequest
func GetAlibabaAlihouseAdminThemeHotUpdateAPIRequest() *AlibabaAlihouseAdminThemeHotUpdateAPIRequest {
	return poolAlibabaAlihouseAdminThemeHotUpdateAPIRequest.Get().(*AlibabaAlihouseAdminThemeHotUpdateAPIRequest)
}

// ReleaseAlibabaAlihouseAdminThemeHotUpdateAPIRequest 将 AlibabaAlihouseAdminThemeHotUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseAdminThemeHotUpdateAPIRequest(v *AlibabaAlihouseAdminThemeHotUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseAdminThemeHotUpdateAPIRequest.Put(v)
}
