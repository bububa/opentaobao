package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseAdminThemeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.admin.theme.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseAdminThemeCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
