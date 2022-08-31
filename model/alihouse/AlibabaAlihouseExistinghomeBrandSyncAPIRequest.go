package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrandSyncAPIRequest 二手房公司品牌数据同步 API请求
// alibaba.alihouse.existinghome.brand.sync
//
// 二手房公司品牌数据同步
type AlibabaAlihouseExistinghomeBrandSyncAPIRequest struct {
	model.Params
	// 入参
	_companyBrandDto *CompanyBrandDto
}

// NewAlibabaAlihouseExistinghomeBrandSyncRequest 初始化AlibabaAlihouseExistinghomeBrandSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeBrandSyncRequest() *AlibabaAlihouseExistinghomeBrandSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeBrandSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeBrandSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.brand.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeBrandSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCompanyBrandDto is CompanyBrandDto Setter
// 入参
func (r *AlibabaAlihouseExistinghomeBrandSyncAPIRequest) SetCompanyBrandDto(_companyBrandDto *CompanyBrandDto) error {
	r._companyBrandDto = _companyBrandDto
	r.Set("company_brand_dto", _companyBrandDto)
	return nil
}

// GetCompanyBrandDto CompanyBrandDto Getter
func (r AlibabaAlihouseExistinghomeBrandSyncAPIRequest) GetCompanyBrandDto() *CompanyBrandDto {
	return r._companyBrandDto
}
