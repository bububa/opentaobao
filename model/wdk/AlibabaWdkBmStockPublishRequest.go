package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销涉及到的商品的库存同步接口 APIRequest
alibaba.wdk.bm.stock.publish

用于操作sku的库存
*/
type AlibabaWdkBmStockPublishRequest struct {
    model.Params

    // 批量入参
    skuStockPublishParamList   []SkuStockPublishParamDo 

}

func NewAlibabaWdkBmStockPublishRequest() *AlibabaWdkBmStockPublishRequest{
    return &AlibabaWdkBmStockPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkBmStockPublishRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.stock.publish"
}

func (r AlibabaWdkBmStockPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkBmStockPublishRequest) SetSkuStockPublishParamList(skuStockPublishParamList []SkuStockPublishParamDo) error {
    r.skuStockPublishParamList = skuStockPublishParamList
    r.Set("sku_stock_publish_param_list", skuStockPublishParamList)
    return nil
}

func (r AlibabaWdkBmStockPublishRequest) GetSkuStockPublishParamList() []SkuStockPublishParamDo {
    return r.skuStockPublishParamList
}

