package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmallitemcentersupplierpriceuploadAPIRequest 天猫服务商服务报价上传 API请求
// tmall.mallitemcenter.supplier.price.upload
//
// 天猫服务商上传服务价格
type TmallmallitemcentersupplierpriceuploadAPIRequest struct {
	model.Params
	// 服务商门店价格列表
	_providerPriceList []StoreOfferPriceDto
	// 服务code
	_serviceCode string
}

// NewTmallmallitemcentersupplierpriceuploadRequest 初始化TmallmallitemcentersupplierpriceuploadAPIRequest对象
func NewTmallmallitemcentersupplierpriceuploadRequest() *TmallmallitemcentersupplierpriceuploadAPIRequest {
	return &TmallmallitemcentersupplierpriceuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmallitemcentersupplierpriceuploadAPIRequest) GetApiMethodName() string {
	return "tmall.mallitemcenter.supplier.price.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmallitemcentersupplierpriceuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmallitemcentersupplierpriceuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProviderPriceList is ProviderPriceList Setter
// 服务商门店价格列表
func (r *TmallmallitemcentersupplierpriceuploadAPIRequest) SetProviderPriceList(_providerPriceList []StoreOfferPriceDto) error {
	r._providerPriceList = _providerPriceList
	r.Set("provider_price_list", _providerPriceList)
	return nil
}

// GetProviderPriceList ProviderPriceList Getter
func (r TmallmallitemcentersupplierpriceuploadAPIRequest) GetProviderPriceList() []StoreOfferPriceDto {
	return r._providerPriceList
}

// SetServiceCode is ServiceCode Setter
// 服务code
func (r *TmallmallitemcentersupplierpriceuploadAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r TmallmallitemcentersupplierpriceuploadAPIRequest) GetServiceCode() string {
	return r._serviceCode
}
