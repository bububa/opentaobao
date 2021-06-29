package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV条码库查询接口 APIRequest
alibaba.lst.pos.open.goods.getgoodsbybarcode

ISV条码库查询接口
*/
type AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest struct {
    model.Params

    // 商品条码
    barcode   string 

}

func NewAlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest() *AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest{
    return &AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.goods.getgoodsbybarcode"
}

func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest) GetBarcode() string {
    return r.barcode
}

