package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务商服务报价上传 APIRequest
tmall.mallitemcenter.supplier.price.upload

天猫服务商上传服务价格
*/
type TmallMallitemcenterSupplierPriceUploadRequest struct {
    model.Params

    // 服务code
    serviceCode   string 

    // 服务商门店价格列表
    providerPriceList   []StoreOfferPriceDto 

}

func NewTmallMallitemcenterSupplierPriceUploadRequest() *TmallMallitemcenterSupplierPriceUploadRequest{
    return &TmallMallitemcenterSupplierPriceUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMallitemcenterSupplierPriceUploadRequest) GetApiMethodName() string {
    return "tmall.mallitemcenter.supplier.price.upload"
}

func (r TmallMallitemcenterSupplierPriceUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMallitemcenterSupplierPriceUploadRequest) SetServiceCode(serviceCode string) error {
    r.serviceCode = serviceCode
    r.Set("service_code", serviceCode)
    return nil
}

func (r TmallMallitemcenterSupplierPriceUploadRequest) GetServiceCode() string {
    return r.serviceCode
}

func (r *TmallMallitemcenterSupplierPriceUploadRequest) SetProviderPriceList(providerPriceList []StoreOfferPriceDto) error {
    r.providerPriceList = providerPriceList
    r.Set("provider_price_list", providerPriceList)
    return nil
}

func (r TmallMallitemcenterSupplierPriceUploadRequest) GetProviderPriceList() []StoreOfferPriceDto {
    return r.providerPriceList
}

