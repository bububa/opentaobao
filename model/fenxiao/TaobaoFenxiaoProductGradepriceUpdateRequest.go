package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据sku设置折扣价 APIRequest
taobao.fenxiao.product.gradeprice.update

供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格
*/
type TaobaoFenxiaoProductGradepriceUpdateRequest struct {
    model.Params

    // 交易类型： AGENT(代销)、DEALER(经销)，ALL(代销和经销)
    tradeType   string 

    // 产品Id
    productId   int64 

    // skuId，如果产品有skuId,必须要输入skuId;没有skuId的时候不必选
    skuId   int64 

    // 选择折扣方式：GRADE（按等级进行设置）;DISCITUTOR(按分销商进行设置）。例如"GRADE,DISTRIBUTOR"
    targetType   string 

    // 会员等级的id或者分销商id，例如：”1001,2001,1002”
    ids   []Number 

    // 优惠价格,大小为0到100000000之间的整数或两位小数，例：优惠价格为：100元2角5分，传入的参数应写成：100.25
    prices   []String 

}

func NewTaobaoFenxiaoProductGradepriceUpdateRequest() *TaobaoFenxiaoProductGradepriceUpdateRequest{
    return &TaobaoFenxiaoProductGradepriceUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductGradepriceUpdateRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.gradeprice.update"
}

func (r TaobaoFenxiaoProductGradepriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductGradepriceUpdateRequest) SetTradeType(tradeType string) error {
    r.tradeType = tradeType
    r.Set("trade_type", tradeType)
    return nil
}

func (r TaobaoFenxiaoProductGradepriceUpdateRequest) GetTradeType() string {
    return r.tradeType
}

func (r *TaobaoFenxiaoProductGradepriceUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoFenxiaoProductGradepriceUpdateRequest) GetProductId() int64 {
    return r.productId
}

func (r *TaobaoFenxiaoProductGradepriceUpdateRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoFenxiaoProductGradepriceUpdateRequest) GetSkuId() int64 {
    return r.skuId
}

func (r *TaobaoFenxiaoProductGradepriceUpdateRequest) SetTargetType(targetType string) error {
    r.targetType = targetType
    r.Set("target_type", targetType)
    return nil
}

func (r TaobaoFenxiaoProductGradepriceUpdateRequest) GetTargetType() string {
    return r.targetType
}

func (r *TaobaoFenxiaoProductGradepriceUpdateRequest) SetIds(ids []Number) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

func (r TaobaoFenxiaoProductGradepriceUpdateRequest) GetIds() []Number {
    return r.ids
}

func (r *TaobaoFenxiaoProductGradepriceUpdateRequest) SetPrices(prices []String) error {
    r.prices = prices
    r.Set("prices", prices)
    return nil
}

func (r TaobaoFenxiaoProductGradepriceUpdateRequest) GetPrices() []String {
    return r.prices
}

