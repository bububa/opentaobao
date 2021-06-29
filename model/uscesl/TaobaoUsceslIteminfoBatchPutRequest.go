package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量写入商品信息接口 APIRequest
taobao.uscesl.iteminfo.batch.put

电子架签批量写入商品数据，用于电子价签展示
*/
type TaobaoUsceslIteminfoBatchPutRequest struct {
    model.Params

    // 商品变更信息集合
    itemChangeBOList   []ItemChangeBo 

    // 门店ID
    shopId   int64 

}

func NewTaobaoUsceslIteminfoBatchPutRequest() *TaobaoUsceslIteminfoBatchPutRequest{
    return &TaobaoUsceslIteminfoBatchPutRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslIteminfoBatchPutRequest) GetApiMethodName() string {
    return "taobao.uscesl.iteminfo.batch.put"
}

func (r TaobaoUsceslIteminfoBatchPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslIteminfoBatchPutRequest) SetItemChangeBOList(itemChangeBOList []ItemChangeBo) error {
    r.itemChangeBOList = itemChangeBOList
    r.Set("item_change_b_o_list", itemChangeBOList)
    return nil
}

func (r TaobaoUsceslIteminfoBatchPutRequest) GetItemChangeBOList() []ItemChangeBo {
    return r.itemChangeBOList
}

func (r *TaobaoUsceslIteminfoBatchPutRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r TaobaoUsceslIteminfoBatchPutRequest) GetShopId() int64 {
    return r.shopId
}

