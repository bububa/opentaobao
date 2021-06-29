package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务商服务报价上传 API请求
tmall.mallitemcenter.supplier.price.upload

天猫服务商上传服务价格
*/
type TmallMallitemcenterSupplierPriceUploadRequest struct {
    model.Params
    // 服务code
    _serviceCode   string
    // 服务商门店价格列表
    _providerPriceList   []StoreOfferPriceDTO
}

// 初始化TmallMallitemcenterSupplierPriceUploadRequest对象
func NewTmallMallitemcenterSupplierPriceUploadRequest() *TmallMallitemcenterSupplierPriceUploadRequest{
    return &TmallMallitemcenterSupplierPriceUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterSupplierPriceUploadRequest) GetApiMethodName() string {
    return "tmall.mallitemcenter.supplier.price.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterSupplierPriceUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceCode Setter
// 服务code
func (r *TmallMallitemcenterSupplierPriceUploadRequest) SetServiceCode(_serviceCode string) error {
    r._serviceCode = _serviceCode
    r.Set("service_code", _serviceCode)
    return nil
}

// ServiceCode Getter
func (r TmallMallitemcenterSupplierPriceUploadRequest) GetServiceCode() string {
    return r._serviceCode
}
// ProviderPriceList Setter
// 服务商门店价格列表
func (r *TmallMallitemcenterSupplierPriceUploadRequest) SetProviderPriceList(_providerPriceList []StoreOfferPriceDTO) error {
    r._providerPriceList = _providerPriceList
    r.Set("provider_price_list", _providerPriceList)
    return nil
}

// ProviderPriceList Getter
func (r TmallMallitemcenterSupplierPriceUploadRequest) GetProviderPriceList() []StoreOfferPriceDTO {
    return r._providerPriceList
}
