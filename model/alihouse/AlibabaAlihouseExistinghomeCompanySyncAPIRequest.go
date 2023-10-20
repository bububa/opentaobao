package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeCompanySyncAPIRequest) Reset() {
	r._companyDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeCompanySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.company.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeCompanySyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeCompanySyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseExistinghomeCompanySyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeCompanySyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeCompanySyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeCompanySyncAPIRequest
func GetAlibabaAlihouseExistinghomeCompanySyncAPIRequest() *AlibabaAlihouseExistinghomeCompanySyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeCompanySyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeCompanySyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeCompanySyncAPIRequest 将 AlibabaAlihouseExistinghomeCompanySyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeCompanySyncAPIRequest(v *AlibabaAlihouseExistinghomeCompanySyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeCompanySyncAPIRequest.Put(v)
}
