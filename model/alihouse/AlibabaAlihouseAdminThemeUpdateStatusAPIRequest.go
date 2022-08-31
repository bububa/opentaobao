package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseAdminThemeUpdateStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.admin.theme.update.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseAdminThemeUpdateStatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
