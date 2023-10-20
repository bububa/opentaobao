package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallMallitemcenterSupplierPriceUploadAPIRequest 天猫服务商服务报价上传 API请求
// tmall.mallitemcenter.supplier.price.upload
//
// 天猫服务商上传服务价格
type TmallMallitemcenterSupplierPriceUploadAPIRequest struct {
	model.Params
	// 服务商门店价格列表
	_providerPriceList []StoreOfferPriceDto
	// 服务code
	_serviceCode string
}

// NewTmallMallitemcenterSupplierPriceUploadRequest 初始化TmallMallitemcenterSupplierPriceUploadAPIRequest对象
func NewTmallMallitemcenterSupplierPriceUploadRequest() *TmallMallitemcenterSupplierPriceUploadAPIRequest {
	return &TmallMallitemcenterSupplierPriceUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterSupplierPriceUploadAPIRequest) GetApiMethodName() string {
	return "tmall.mallitemcenter.supplier.price.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterSupplierPriceUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMallitemcenterSupplierPriceUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProviderPriceList is ProviderPriceList Setter
// 服务商门店价格列表
func (r *TmallMallitemcenterSupplierPriceUploadAPIRequest) SetProviderPriceList(_providerPriceList []StoreOfferPriceDto) error {
	r._providerPriceList = _providerPriceList
	r.Set("provider_price_list", _providerPriceList)
	return nil
}

// GetProviderPriceList ProviderPriceList Getter
func (r TmallMallitemcenterSupplierPriceUploadAPIRequest) GetProviderPriceList() []StoreOfferPriceDto {
	return r._providerPriceList
}

// SetServiceCode is ServiceCode Setter
// 服务code
func (r *TmallMallitemcenterSupplierPriceUploadAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r TmallMallitemcenterSupplierPriceUploadAPIRequest) GetServiceCode() string {
	return r._serviceCode
}
