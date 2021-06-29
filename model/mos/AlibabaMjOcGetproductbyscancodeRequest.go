package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
POS商品查询接口 API请求
alibaba.mj.oc.getproductbyscancode

此API用于在银泰商场中，POS端扫码获取商品信息
*/
type AlibabaMjOcGetproductbyscancodeRequest struct {
    model.Params
    // 码, 对应的信息可能是款号，也有可能是具体的某一个商品
    code   string
    // 条码/二维码/rfid(电子标签),货号、条码、零售+唯一码;ARTNO、BARCODE、UNIQUECODE
    codeType   string
    // 专柜编码
    shopCode   string
    // 门店编码
    storeCode   string
}

// 初始化AlibabaMjOcGetproductbyscancodeRequest对象
func NewAlibabaMjOcGetproductbyscancodeRequest() *AlibabaMjOcGetproductbyscancodeRequest{
    return &AlibabaMjOcGetproductbyscancodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcGetproductbyscancodeRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.getproductbyscancode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcGetproductbyscancodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 码, 对应的信息可能是款号，也有可能是具体的某一个商品
func (r *AlibabaMjOcGetproductbyscancodeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaMjOcGetproductbyscancodeRequest) GetCode() string {
    return r.code
}
// CodeType Setter
// 条码/二维码/rfid(电子标签),货号、条码、零售+唯一码;ARTNO、BARCODE、UNIQUECODE
func (r *AlibabaMjOcGetproductbyscancodeRequest) SetCodeType(codeType string) error {
    r.codeType = codeType
    r.Set("code_type", codeType)
    return nil
}

// CodeType Getter
func (r AlibabaMjOcGetproductbyscancodeRequest) GetCodeType() string {
    return r.codeType
}
// ShopCode Setter
// 专柜编码
func (r *AlibabaMjOcGetproductbyscancodeRequest) SetShopCode(shopCode string) error {
    r.shopCode = shopCode
    r.Set("shop_code", shopCode)
    return nil
}

// ShopCode Getter
func (r AlibabaMjOcGetproductbyscancodeRequest) GetShopCode() string {
    return r.shopCode
}
// StoreCode Setter
// 门店编码
func (r *AlibabaMjOcGetproductbyscancodeRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r AlibabaMjOcGetproductbyscancodeRequest) GetStoreCode() string {
    return r.storeCode
}
