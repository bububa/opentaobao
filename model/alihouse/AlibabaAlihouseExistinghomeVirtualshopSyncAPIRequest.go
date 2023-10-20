package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest 二手房虚拟店铺数据同步 API请求
// alibaba.alihouse.existinghome.virtualshop.sync
//
// 二手房虚拟店铺数据同步
type AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest struct {
	model.Params
	// 入参
	_companyVirtualShopDto *CompanyVirtualShopDto
}

// NewAlibabaAlihouseExistinghomeVirtualshopSyncRequest 初始化AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeVirtualshopSyncRequest() *AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.virtualshop.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyVirtualShopDto is CompanyVirtualShopDto Setter
// 入参
func (r *AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest) SetCompanyVirtualShopDto(_companyVirtualShopDto *CompanyVirtualShopDto) error {
	r._companyVirtualShopDto = _companyVirtualShopDto
	r.Set("company_virtual_shop_dto", _companyVirtualShopDto)
	return nil
}

// GetCompanyVirtualShopDto CompanyVirtualShopDto Getter
func (r AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest) GetCompanyVirtualShopDto() *CompanyVirtualShopDto {
	return r._companyVirtualShopDto
}
