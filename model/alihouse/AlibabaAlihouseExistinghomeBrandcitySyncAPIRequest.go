package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest) Reset() {
	r._companyBrandCityDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.brandcity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseExistinghomeBrandcitySyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeBrandcitySyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeBrandcitySyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest
func GetAlibabaAlihouseExistinghomeBrandcitySyncAPIRequest() *AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeBrandcitySyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeBrandcitySyncAPIRequest 将 AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBrandcitySyncAPIRequest(v *AlibabaAlihouseExistinghomeBrandcitySyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBrandcitySyncAPIRequest.Put(v)
}
