package fivee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
进口商品发布 APIRequest
taobao.fivee.importproduct.publish

直营业务商家入住发布商品时，上传商品及商家证照信息
*/
type TaobaoFiveeImportproductPublishRequest struct {
    model.Params

    // 进口商品
    importProduct   *ImportProduct 

    // bu身份标识
    paramBucode   string 

}

func NewTaobaoFiveeImportproductPublishRequest() *TaobaoFiveeImportproductPublishRequest{
    return &TaobaoFiveeImportproductPublishRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFiveeImportproductPublishRequest) GetApiMethodName() string {
    return "taobao.fivee.importproduct.publish"
}

func (r TaobaoFiveeImportproductPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFiveeImportproductPublishRequest) SetImportProduct(importProduct *ImportProduct) error {
    r.importProduct = importProduct
    r.Set("import_product", importProduct)
    return nil
}

func (r TaobaoFiveeImportproductPublishRequest) GetImportProduct() *ImportProduct {
    return r.importProduct
}

func (r *TaobaoFiveeImportproductPublishRequest) SetParamBucode(paramBucode string) error {
    r.paramBucode = paramBucode
    r.Set("param_bucode", paramBucode)
    return nil
}

func (r TaobaoFiveeImportproductPublishRequest) GetParamBucode() string {
    return r.paramBucode
}

