package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按商家批量写入商品接口 APIRequest
taobao.uscesl.iteminfo.batch.insert

【电子价签】支持按照商家-门店维度批量写入商品数据
*/
type TaobaoUsceslIteminfoBatchInsertRequest struct {
    model.Params

    // 商品信息集合，限制500
    itemList   []ItemChangeBo 

    // 门店ID
    storeId   int64 

    // 商家KEY
    bizBrandKey   string 

}

func NewTaobaoUsceslIteminfoBatchInsertRequest() *TaobaoUsceslIteminfoBatchInsertRequest{
    return &TaobaoUsceslIteminfoBatchInsertRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslIteminfoBatchInsertRequest) GetApiMethodName() string {
    return "taobao.uscesl.iteminfo.batch.insert"
}

func (r TaobaoUsceslIteminfoBatchInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslIteminfoBatchInsertRequest) SetItemList(itemList []ItemChangeBo) error {
    r.itemList = itemList
    r.Set("item_list", itemList)
    return nil
}

func (r TaobaoUsceslIteminfoBatchInsertRequest) GetItemList() []ItemChangeBo {
    return r.itemList
}

func (r *TaobaoUsceslIteminfoBatchInsertRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoUsceslIteminfoBatchInsertRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoUsceslIteminfoBatchInsertRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

func (r TaobaoUsceslIteminfoBatchInsertRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}

