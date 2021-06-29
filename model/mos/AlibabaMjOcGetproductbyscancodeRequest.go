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
    _code   string
    // 条码/二维码/rfid(电子标签),货号、条码、零售+唯一码;ARTNO、BARCODE、UNIQUECODE
    _codeType   string
    // 专柜编码
    _shopCode   string
    // 门店编码
    _storeCode   string
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
func (r *AlibabaMjOcGetproductbyscancodeRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaMjOcGetproductbyscancodeRequest) GetCode() string {
    return r._code
}
// CodeType Setter
// 条码/二维码/rfid(电子标签),货号、条码、零售+唯一码;ARTNO、BARCODE、UNIQUECODE
func (r *AlibabaMjOcGetproductbyscancodeRequest) SetCodeType(_codeType string) error {
    r._codeType = _codeType
    r.Set("code_type", _codeType)
    return nil
}

// CodeType Getter
func (r AlibabaMjOcGetproductbyscancodeRequest) GetCodeType() string {
    return r._codeType
}
// ShopCode Setter
// 专柜编码
func (r *AlibabaMjOcGetproductbyscancodeRequest) SetShopCode(_shopCode string) error {
    r._shopCode = _shopCode
    r.Set("shop_code", _shopCode)
    return nil
}

// ShopCode Getter
func (r AlibabaMjOcGetproductbyscancodeRequest) GetShopCode() string {
    return r._shopCode
}
// StoreCode Setter
// 门店编码
func (r *AlibabaMjOcGetproductbyscancodeRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r AlibabaMjOcGetproductbyscancodeRequest) GetStoreCode() string {
    return r._storeCode
}
