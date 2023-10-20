package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebrandsyncAPIRequest 二手房公司品牌数据同步 API请求
// alibaba.alihouse.existinghome.brand.sync
//
// 二手房公司品牌数据同步
type AlibabaalihouseexistinghomebrandsyncAPIRequest struct {
	model.Params
	// 入参
	_companyBrandDto *CompanyBrandDto
}

// NewAlibabaalihouseexistinghomebrandsyncRequest 初始化AlibabaalihouseexistinghomebrandsyncAPIRequest对象
func NewAlibabaalihouseexistinghomebrandsyncRequest() *AlibabaalihouseexistinghomebrandsyncAPIRequest {
	return &AlibabaalihouseexistinghomebrandsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomebrandsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.brand.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomebrandsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomebrandsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyBrandDto is CompanyBrandDto Setter
// 入参
func (r *AlibabaalihouseexistinghomebrandsyncAPIRequest) SetCompanyBrandDto(_companyBrandDto *CompanyBrandDto) error {
	r._companyBrandDto = _companyBrandDto
	r.Set("company_brand_dto", _companyBrandDto)
	return nil
}

// GetCompanyBrandDto CompanyBrandDto Getter
func (r AlibabaalihouseexistinghomebrandsyncAPIRequest) GetCompanyBrandDto() *CompanyBrandDto {
	return r._companyBrandDto
}
