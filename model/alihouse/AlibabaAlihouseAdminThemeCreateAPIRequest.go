package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseadminthemecreateAPIRequest 创建云主题 API请求
// alibaba.alihouse.admin.theme.create
//
// 创建云主题
type AlibabaalihouseadminthemecreateAPIRequest struct {
	model.Params
	// 请求云主题参数
	_etcThemeDto *EtcThemeDto
}

// NewAlibabaalihouseadminthemecreateRequest 初始化AlibabaalihouseadminthemecreateAPIRequest对象
func NewAlibabaalihouseadminthemecreateRequest() *AlibabaalihouseadminthemecreateAPIRequest {
	return &AlibabaalihouseadminthemecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseadminthemecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.admin.theme.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseadminthemecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseadminthemecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEtcThemeDto is EtcThemeDto Setter
// 请求云主题参数
func (r *AlibabaalihouseadminthemecreateAPIRequest) SetEtcThemeDto(_etcThemeDto *EtcThemeDto) error {
	r._etcThemeDto = _etcThemeDto
	r.Set("etc_theme_dto", _etcThemeDto)
	return nil
}

// GetEtcThemeDto EtcThemeDto Getter
func (r AlibabaalihouseadminthemecreateAPIRequest) GetEtcThemeDto() *EtcThemeDto {
	return r._etcThemeDto
}
