package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV条码库查询接口 API请求
alibaba.lst.pos.open.goods.getgoodsbybarcode

ISV条码库查询接口
*/
type AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest struct {
    model.Params
    // 商品条码
    barcode   string
}

// 初始化AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest对象
func NewAlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest() *AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest{
    return &AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.goods.getgoodsbybarcode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Barcode Setter
// 商品条码
func (r *AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

// Barcode Getter
func (r AlibabaLstPosOpenGoodsGetgoodsbybarcodeRequest) GetBarcode() string {
    return r.barcode
}
