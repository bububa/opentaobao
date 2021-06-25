package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
派样商品门店库存查询接口 APIRequest
alibaba.wdk.bm.paiyang.stock.query

淘鲜达接入第三方进行派样，第三方查询派样商品的门店库存信息。
*/
type AlibabaWdkBmPaiyangStockQueryRequest struct {
    model.Params

    // 请求入参
    isvShopStockParam   *IsvShopStockParam 

}

func NewAlibabaWdkBmPaiyangStockQueryRequest() *AlibabaWdkBmPaiyangStockQueryRequest{
    return &AlibabaWdkBmPaiyangStockQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkBmPaiyangStockQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.paiyang.stock.query"
}

func (r AlibabaWdkBmPaiyangStockQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkBmPaiyangStockQueryRequest) SetIsvShopStockParam(isvShopStockParam *IsvShopStockParam) error {
    r.isvShopStockParam = isvShopStockParam
    r.Set("isv_shop_stock_param", isvShopStockParam)
    return nil
}

func (r AlibabaWdkBmPaiyangStockQueryRequest) GetIsvShopStockParam() *IsvShopStockParam {
    return r.isvShopStockParam
}

