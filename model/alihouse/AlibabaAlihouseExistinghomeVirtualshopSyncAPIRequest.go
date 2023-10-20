package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomevirtualshopsyncAPIRequest 二手房虚拟店铺数据同步 API请求
// alibaba.alihouse.existinghome.virtualshop.sync
//
// 二手房虚拟店铺数据同步
type AlibabaalihouseexistinghomevirtualshopsyncAPIRequest struct {
	model.Params
	// 入参
	_companyVirtualShopDto *CompanyVirtualShopDto
}

// NewAlibabaalihouseexistinghomevirtualshopsyncRequest 初始化AlibabaalihouseexistinghomevirtualshopsyncAPIRequest对象
func NewAlibabaalihouseexistinghomevirtualshopsyncRequest() *AlibabaalihouseexistinghomevirtualshopsyncAPIRequest {
	return &AlibabaalihouseexistinghomevirtualshopsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomevirtualshopsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.virtualshop.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomevirtualshopsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomevirtualshopsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyVirtualShopDto is CompanyVirtualShopDto Setter
// 入参
func (r *AlibabaalihouseexistinghomevirtualshopsyncAPIRequest) SetCompanyVirtualShopDto(_companyVirtualShopDto *CompanyVirtualShopDto) error {
	r._companyVirtualShopDto = _companyVirtualShopDto
	r.Set("company_virtual_shop_dto", _companyVirtualShopDto)
	return nil
}

// GetCompanyVirtualShopDto CompanyVirtualShopDto Getter
func (r AlibabaalihouseexistinghomevirtualshopsyncAPIRequest) GetCompanyVirtualShopDto() *CompanyVirtualShopDto {
	return r._companyVirtualShopDto
}
