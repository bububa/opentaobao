package fivee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
进口商品查询 APIRequest
taobao.fivee.importproduct.get

资质共享平台查询进口商品信息
*/
type TaobaoFiveeImportproductGetRequest struct {
    model.Params

    // bu身份标识
    paramBuCode   string 

    // 条形码
    paramBarcode   string 

}

func NewTaobaoFiveeImportproductGetRequest() *TaobaoFiveeImportproductGetRequest{
    return &TaobaoFiveeImportproductGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFiveeImportproductGetRequest) GetApiMethodName() string {
    return "taobao.fivee.importproduct.get"
}

func (r TaobaoFiveeImportproductGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFiveeImportproductGetRequest) SetParamBuCode(paramBuCode string) error {
    r.paramBuCode = paramBuCode
    r.Set("param_bu_code", paramBuCode)
    return nil
}

func (r TaobaoFiveeImportproductGetRequest) GetParamBuCode() string {
    return r.paramBuCode
}

func (r *TaobaoFiveeImportproductGetRequest) SetParamBarcode(paramBarcode string) error {
    r.paramBarcode = paramBarcode
    r.Set("param_barcode", paramBarcode)
    return nil
}

func (r TaobaoFiveeImportproductGetRequest) GetParamBarcode() string {
    return r.paramBarcode
}

