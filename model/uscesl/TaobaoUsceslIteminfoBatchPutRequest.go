package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量写入商品信息接口 API请求
taobao.uscesl.iteminfo.batch.put

电子架签批量写入商品数据，用于电子价签展示
*/
type TaobaoUsceslIteminfoBatchPutRequest struct {
    model.Params
    // 商品变更信息集合
    _itemChangeBOList   []ItemChangeBo
    // 门店ID
    _shopId   int64
}

// 初始化TaobaoUsceslIteminfoBatchPutRequest对象
func NewTaobaoUsceslIteminfoBatchPutRequest() *TaobaoUsceslIteminfoBatchPutRequest{
    return &TaobaoUsceslIteminfoBatchPutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslIteminfoBatchPutRequest) GetApiMethodName() string {
    return "taobao.uscesl.iteminfo.batch.put"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslIteminfoBatchPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemChangeBOList Setter
// 商品变更信息集合
func (r *TaobaoUsceslIteminfoBatchPutRequest) SetItemChangeBOList(_itemChangeBOList []ItemChangeBo) error {
    r._itemChangeBOList = _itemChangeBOList
    r.Set("item_change_b_o_list", _itemChangeBOList)
    return nil
}

// ItemChangeBOList Getter
func (r TaobaoUsceslIteminfoBatchPutRequest) GetItemChangeBOList() []ItemChangeBo {
    return r._itemChangeBOList
}
// ShopId Setter
// 门店ID
func (r *TaobaoUsceslIteminfoBatchPutRequest) SetShopId(_shopId int64) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r TaobaoUsceslIteminfoBatchPutRequest) GetShopId() int64 {
    return r._shopId
}
