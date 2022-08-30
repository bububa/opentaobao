package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest 二手房城市品牌店数据同步 API请求
// alibaba.alihouse.existinghome.brandcity.sync
//
// 二手房城市品牌店数据同步
type AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest struct {
	model.Params
	// 入参
	_companyBrandCityDto *CompanyBrandCityDto
}

// NewAlibabaAlihouseExistinghomeBrandcitySyncRequest 初始化AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeBrandcitySyncRequest() *AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest {
	return &AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.brandcity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCompanyBrandCityDto is CompanyBrandCityDto Setter
// 入参
func (r *AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest) SetCompanyBrandCityDto(_companyBrandCityDto *CompanyBrandCityDto) error {
	r._companyBrandCityDto = _companyBrandCityDto
	r.Set("company_brand_city_dto", _companyBrandCityDto)
	return nil
}

// GetCompanyBrandCityDto CompanyBrandCityDto Getter
func (r AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest) GetCompanyBrandCityDto() *CompanyBrandCityDto {
	return r._companyBrandCityDto
}
