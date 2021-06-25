package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商品条码查询接口 
alibaba.wdk.sku.barcode.query

查询商品编码，支持一品多码
*/
func AlibabaWdkSkuBarcodeQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuBarcodeQueryRequest, session string) (*wdk.AlibabaWdkSkuBarcodeQueryResponse, error) {
    var resp wdk.AlibabaWdkSkuBarcodeQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
