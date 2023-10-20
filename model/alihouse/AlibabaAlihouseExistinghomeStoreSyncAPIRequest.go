package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomestoresyncAPIRequest 二手房标准门店数据同步 API请求
// alibaba.alihouse.existinghome.store.sync
//
// 二手房标准门店数据同步
type AlibabaalihouseexistinghomestoresyncAPIRequest struct {
	model.Params
	// 入参
	_companyStoreDto *CompanyStoreDto
}

// NewAlibabaalihouseexistinghomestoresyncRequest 初始化AlibabaalihouseexistinghomestoresyncAPIRequest对象
func NewAlibabaalihouseexistinghomestoresyncRequest() *AlibabaalihouseexistinghomestoresyncAPIRequest {
	return &AlibabaalihouseexistinghomestoresyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomestoresyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.store.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomestoresyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomestoresyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyStoreDto is CompanyStoreDto Setter
// 入参
func (r *AlibabaalihouseexistinghomestoresyncAPIRequest) SetCompanyStoreDto(_companyStoreDto *CompanyStoreDto) error {
	r._companyStoreDto = _companyStoreDto
	r.Set("company_store_dto", _companyStoreDto)
	return nil
}

// GetCompanyStoreDto CompanyStoreDto Getter
func (r AlibabaalihouseexistinghomestoresyncAPIRequest) GetCompanyStoreDto() *CompanyStoreDto {
	return r._companyStoreDto
}
