package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
更新商品条形码信息 
taobao.item.barcode.update

通过该接口，将商品以及SKU上得条形码信息补全
*/
func TaobaoItemBarcodeUpdate(clt *core.SDKClient, req *product.TaobaoItemBarcodeUpdateRequest, session string) (*product.TaobaoItemBarcodeUpdateAPIResponse, error) {
    var resp product.TaobaoItemBarcodeUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
