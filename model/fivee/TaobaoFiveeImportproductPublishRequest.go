package fivee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
进口商品发布 API请求
taobao.fivee.importproduct.publish

直营业务商家入住发布商品时，上传商品及商家证照信息
*/
type TaobaoFiveeImportproductPublishRequest struct {
    model.Params
    // 进口商品
    _importProduct   *ImportProduct
    // bu身份标识
    _paramBucode   string
}

// 初始化TaobaoFiveeImportproductPublishRequest对象
func NewTaobaoFiveeImportproductPublishRequest() *TaobaoFiveeImportproductPublishRequest{
    return &TaobaoFiveeImportproductPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFiveeImportproductPublishRequest) GetApiMethodName() string {
    return "taobao.fivee.importproduct.publish"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFiveeImportproductPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImportProduct Setter
// 进口商品
func (r *TaobaoFiveeImportproductPublishRequest) SetImportProduct(_importProduct *ImportProduct) error {
    r._importProduct = _importProduct
    r.Set("import_product", _importProduct)
    return nil
}

// ImportProduct Getter
func (r TaobaoFiveeImportproductPublishRequest) GetImportProduct() *ImportProduct {
    return r._importProduct
}
// ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeImportproductPublishRequest) SetParamBucode(_paramBucode string) error {
    r._paramBucode = _paramBucode
    r.Set("param_bucode", _paramBucode)
    return nil
}

// ParamBucode Getter
func (r TaobaoFiveeImportproductPublishRequest) GetParamBucode() string {
    return r._paramBucode
}