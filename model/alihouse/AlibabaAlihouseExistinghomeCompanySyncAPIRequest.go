package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomecompanysyncAPIRequest 二手房公司同步接口 API请求
// alibaba.alihouse.existinghome.company.sync
//
// 二手房公司同步接口
type AlibabaalihouseexistinghomecompanysyncAPIRequest struct {
	model.Params
	// 入参
	_companyDto *CompanyDto
}

// NewAlibabaalihouseexistinghomecompanysyncRequest 初始化AlibabaalihouseexistinghomecompanysyncAPIRequest对象
func NewAlibabaalihouseexistinghomecompanysyncRequest() *AlibabaalihouseexistinghomecompanysyncAPIRequest {
	return &AlibabaalihouseexistinghomecompanysyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomecompanysyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.company.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomecompanysyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomecompanysyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyDto is CompanyDto Setter
// 入参
func (r *AlibabaalihouseexistinghomecompanysyncAPIRequest) SetCompanyDto(_companyDto *CompanyDto) error {
	r._companyDto = _companyDto
	r.Set("company_dto", _companyDto)
	return nil
}

// GetCompanyDto CompanyDto Getter
func (r AlibabaalihouseexistinghomecompanysyncAPIRequest) GetCompanyDto() *CompanyDto {
	return r._companyDto
}
