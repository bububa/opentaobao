package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebrandcitysyncAPIRequest 二手房城市品牌店数据同步 API请求
// alibaba.alihouse.existinghome.brandcity.sync
//
// 二手房城市品牌店数据同步
type AlibabaalihouseexistinghomebrandcitysyncAPIRequest struct {
	model.Params
	// 入参
	_companyBrandCityDto *CompanyBrandCityDto
}

// NewAlibabaalihouseexistinghomebrandcitysyncRequest 初始化AlibabaalihouseexistinghomebrandcitysyncAPIRequest对象
func NewAlibabaalihouseexistinghomebrandcitysyncRequest() *AlibabaalihouseexistinghomebrandcitysyncAPIRequest {
	return &AlibabaalihouseexistinghomebrandcitysyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomebrandcitysyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.brandcity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomebrandcitysyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomebrandcitysyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyBrandCityDto is CompanyBrandCityDto Setter
// 入参
func (r *AlibabaalihouseexistinghomebrandcitysyncAPIRequest) SetCompanyBrandCityDto(_companyBrandCityDto *CompanyBrandCityDto) error {
	r._companyBrandCityDto = _companyBrandCityDto
	r.Set("company_brand_city_dto", _companyBrandCityDto)
	return nil
}

// GetCompanyBrandCityDto CompanyBrandCityDto Getter
func (r AlibabaalihouseexistinghomebrandcitysyncAPIRequest) GetCompanyBrandCityDto() *CompanyBrandCityDto {
	return r._companyBrandCityDto
}
