package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销涉及到的商品的库存同步接口 API请求
alibaba.wdk.bm.stock.publish

用于操作sku的库存
*/
type AlibabaWdkBmStockPublishRequest struct {
    model.Params
    // 批量入参
    _skuStockPublishParamList   []SkuStockPublishParamDO
}

// 初始化AlibabaWdkBmStockPublishRequest对象
func NewAlibabaWdkBmStockPublishRequest() *AlibabaWdkBmStockPublishRequest{
    return &AlibabaWdkBmStockPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmStockPublishRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.stock.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmStockPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuStockPublishParamList Setter
// 批量入参
func (r *AlibabaWdkBmStockPublishRequest) SetSkuStockPublishParamList(_skuStockPublishParamList []SkuStockPublishParamDO) error {
    r._skuStockPublishParamList = _skuStockPublishParamList
    r.Set("sku_stock_publish_param_list", _skuStockPublishParamList)
    return nil
}

// SkuStockPublishParamList Getter
func (r AlibabaWdkBmStockPublishRequest) GetSkuStockPublishParamList() []SkuStockPublishParamDO {
    return r._skuStockPublishParamList
}
