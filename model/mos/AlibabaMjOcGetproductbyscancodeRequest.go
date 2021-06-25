package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
POS商品查询接口 APIRequest
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

func NewAlibabaMjOcGetproductbyscancodeRequest() *AlibabaMjOcGetproductbyscancodeRequest{
    return &AlibabaMjOcGetproductbyscancodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjOcGetproductbyscancodeRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.getproductbyscancode"
}

func (r AlibabaMjOcGetproductbyscancodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjOcGetproductbyscancodeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaMjOcGetproductbyscancodeRequest) GetCode() string {
    return r.code
}

func (r *AlibabaMjOcGetproductbyscancodeRequest) SetCodeType(codeType string) error {
    r.codeType = codeType
    r.Set("code_type", codeType)
    return nil
}

func (r AlibabaMjOcGetproductbyscancodeRequest) GetCodeType() string {
    return r.codeType
}

func (r *AlibabaMjOcGetproductbyscancodeRequest) SetShopCode(shopCode string) error {
    r.shopCode = shopCode
    r.Set("shop_code", shopCode)
    return nil
}

func (r AlibabaMjOcGetproductbyscancodeRequest) GetShopCode() string {
    return r.shopCode
}

func (r *AlibabaMjOcGetproductbyscancodeRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r AlibabaMjOcGetproductbyscancodeRequest) GetStoreCode() string {
    return r.storeCode
}

