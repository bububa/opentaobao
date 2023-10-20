package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseadminthemeupdateAPIRequest 云主题更新 API请求
// alibaba.alihouse.admin.theme.update
//
// 云主题更新
type AlibabaalihouseadminthemeupdateAPIRequest struct {
	model.Params
	// 请求云主题参数
	_etcThemeDto *EtcThemeDto
}

// NewAlibabaalihouseadminthemeupdateRequest 初始化AlibabaalihouseadminthemeupdateAPIRequest对象
func NewAlibabaalihouseadminthemeupdateRequest() *AlibabaalihouseadminthemeupdateAPIRequest {
	return &AlibabaalihouseadminthemeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseadminthemeupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.admin.theme.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseadminthemeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseadminthemeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEtcThemeDto is EtcThemeDto Setter
// 请求云主题参数
func (r *AlibabaalihouseadminthemeupdateAPIRequest) SetEtcThemeDto(_etcThemeDto *EtcThemeDto) error {
	r._etcThemeDto = _etcThemeDto
	r.Set("etc_theme_dto", _etcThemeDto)
	return nil
}

// GetEtcThemeDto EtcThemeDto Getter
func (r AlibabaalihouseadminthemeupdateAPIRequest) GetEtcThemeDto() *EtcThemeDto {
	return r._etcThemeDto
}
