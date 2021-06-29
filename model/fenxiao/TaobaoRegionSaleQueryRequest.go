package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品销售区域 API请求
taobao.region.sale.query

查询商品销售区域
*/
type TaobaoRegionSaleQueryRequest struct {
    model.Params
    // 商品id
    itemId   int64
    // 无sku传0
    skuId   int64
    // 1:国家;2:省;3: 市;4:区县
    saleRegionLevel   int64
}

// 初始化TaobaoRegionSaleQueryRequest对象
func NewTaobaoRegionSaleQueryRequest() *TaobaoRegionSaleQueryRequest{
    return &TaobaoRegionSaleQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRegionSaleQueryRequest) GetApiMethodName() string {
    return "taobao.region.sale.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRegionSaleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoRegionSaleQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoRegionSaleQueryRequest) GetItemId() int64 {
    return r.itemId
}
// SkuId Setter
// 无sku传0
func (r *TaobaoRegionSaleQueryRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r TaobaoRegionSaleQueryRequest) GetSkuId() int64 {
    return r.skuId
}
// SaleRegionLevel Setter
// 1:国家;2:省;3: 市;4:区县
func (r *TaobaoRegionSaleQueryRequest) SetSaleRegionLevel(saleRegionLevel int64) error {
    r.saleRegionLevel = saleRegionLevel
    r.Set("sale_region_level", saleRegionLevel)
    return nil
}

// SaleRegionLevel Getter
func (r TaobaoRegionSaleQueryRequest) GetSaleRegionLevel() int64 {
    return r.saleRegionLevel
}
