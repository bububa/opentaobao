package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalsolutionserviceresourcequeryAPIRequest 查询解决方案服务资源列表 API请求
// cainiao.global.solution.service.resource.query
//
// 返回直接解决方案的指定物流服务的可用资源列表
type CainiaoglobalsolutionserviceresourcequeryAPIRequest struct {
	model.Params
	// 多语言信息
	_locale string
	// 查询参数
	_solutionServiceResParam *QuerySolutionServiceResParam
	// 发件信息
	_senderParam *OpenSenderParam
}

// NewCainiaoglobalsolutionserviceresourcequeryRequest 初始化CainiaoglobalsolutionserviceresourcequeryAPIRequest对象
func NewCainiaoglobalsolutionserviceresourcequeryRequest() *CainiaoglobalsolutionserviceresourcequeryAPIRequest {
	return &CainiaoglobalsolutionserviceresourcequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalsolutionserviceresourcequeryAPIRequest) GetApiMethodName() string {
	return "cainiao.global.solution.service.resource.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalsolutionserviceresourcequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalsolutionserviceresourcequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocale is Locale Setter
// 多语言信息
func (r *CainiaoglobalsolutionserviceresourcequeryAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoglobalsolutionserviceresourcequeryAPIRequest) GetLocale() string {
	return r._locale
}

// SetSolutionServiceResParam is SolutionServiceResParam Setter
// 查询参数
func (r *CainiaoglobalsolutionserviceresourcequeryAPIRequest) SetSolutionServiceResParam(_solutionServiceResParam *QuerySolutionServiceResParam) error {
	r._solutionServiceResParam = _solutionServiceResParam
	r.Set("solution_service_res_param", _solutionServiceResParam)
	return nil
}

// GetSolutionServiceResParam SolutionServiceResParam Getter
func (r CainiaoglobalsolutionserviceresourcequeryAPIRequest) GetSolutionServiceResParam() *QuerySolutionServiceResParam {
	return r._solutionServiceResParam
}

// SetSenderParam is SenderParam Setter
// 发件信息
func (r *CainiaoglobalsolutionserviceresourcequeryAPIRequest) SetSenderParam(_senderParam *OpenSenderParam) error {
	r._senderParam = _senderParam
	r.Set("sender_param", _senderParam)
	return nil
}

// GetSenderParam SenderParam Getter
func (r CainiaoglobalsolutionserviceresourcequeryAPIRequest) GetSenderParam() *OpenSenderParam {
	return r._senderParam
}
