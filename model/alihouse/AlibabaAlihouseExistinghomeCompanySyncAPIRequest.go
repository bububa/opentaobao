package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeCompanySyncAPIRequest 二手房公司同步接口 API请求
// alibaba.alihouse.existinghome.company.sync
//
// 二手房公司同步接口
type AlibabaAlihouseExistinghomeCompanySyncAPIRequest struct {
	model.Params
	// 入参
	_companyDto *CompanyDto
}

// NewAlibabaAlihouseExistinghomeCompanySyncRequest 初始化AlibabaAlihouseExistinghomeCompanySyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeCompanySyncRequest() *AlibabaAlihouseExistinghomeCompanySyncAPIRequest {
	return &AlibabaAlihouseExistinghomeCompanySyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeCompanySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.company.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeCompanySyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCompanyDto is CompanyDto Setter
// 入参
func (r *AlibabaAlihouseExistinghomeCompanySyncAPIRequest) SetCompanyDto(_companyDto *CompanyDto) error {
	r._companyDto = _companyDto
	r.Set("company_dto", _companyDto)
	return nil
}

// GetCompanyDto CompanyDto Getter
func (r AlibabaAlihouseExistinghomeCompanySyncAPIRequest) GetCompanyDto() *CompanyDto {
	return r._companyDto
}
