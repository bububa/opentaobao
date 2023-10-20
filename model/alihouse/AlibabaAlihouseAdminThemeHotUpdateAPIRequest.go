package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseadminthemehotupdateAPIRequest 云主题热更新数据集 API请求
// alibaba.alihouse.admin.theme.hot.update
//
// 云主题更新
type AlibabaalihouseadminthemehotupdateAPIRequest struct {
	model.Params
	// 请求云主题参数
	_etcThemeDto *EtcThemeDto
}

// NewAlibabaalihouseadminthemehotupdateRequest 初始化AlibabaalihouseadminthemehotupdateAPIRequest对象
func NewAlibabaalihouseadminthemehotupdateRequest() *AlibabaalihouseadminthemehotupdateAPIRequest {
	return &AlibabaalihouseadminthemehotupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseadminthemehotupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.admin.theme.hot.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseadminthemehotupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseadminthemehotupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEtcThemeDto is EtcThemeDto Setter
// 请求云主题参数
func (r *AlibabaalihouseadminthemehotupdateAPIRequest) SetEtcThemeDto(_etcThemeDto *EtcThemeDto) error {
	r._etcThemeDto = _etcThemeDto
	r.Set("etc_theme_dto", _etcThemeDto)
	return nil
}

// GetEtcThemeDto EtcThemeDto Getter
func (r AlibabaalihouseadminthemehotupdateAPIRequest) GetEtcThemeDto() *EtcThemeDto {
	return r._etcThemeDto
}
