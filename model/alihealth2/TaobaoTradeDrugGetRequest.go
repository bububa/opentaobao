package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家未确认订单列表 APIRequest
taobao.trade.drug.get

可以按商家或是店铺维度的进行查询买家付款卖家未确认订单，一次返回不大于20条订单
*/
type TaobaoTradeDrugGetRequest struct {
    model.Params

    // 店铺id
    storeId   int64 

    // true-商家下所有店铺的待确认订单, false—指定店铺的订单
    isAll   bool 

    // 返回记录数，超过20按20条返回数据
    maxSize   int64 

}

func NewTaobaoTradeDrugGetRequest() *TaobaoTradeDrugGetRequest{
    return &TaobaoTradeDrugGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeDrugGetRequest) GetApiMethodName() string {
    return "taobao.trade.drug.get"
}

func (r TaobaoTradeDrugGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeDrugGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoTradeDrugGetRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoTradeDrugGetRequest) SetIsAll(isAll bool) error {
    r.isAll = isAll
    r.Set("is_all", isAll)
    return nil
}

func (r TaobaoTradeDrugGetRequest) GetIsAll() bool {
    return r.isAll
}

func (r *TaobaoTradeDrugGetRequest) SetMaxSize(maxSize int64) error {
    r.maxSize = maxSize
    r.Set("max_size", maxSize)
    return nil
}

func (r TaobaoTradeDrugGetRequest) GetMaxSize() int64 {
    return r.maxSize
}

