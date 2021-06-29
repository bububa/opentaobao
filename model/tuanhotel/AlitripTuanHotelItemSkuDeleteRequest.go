package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店团购套餐商品SKU删除 API请求
alitrip.tuan.hotel.item.sku.delete

商户对发布的宝贝套餐价格库存信息进行删除
*/
type AlitripTuanHotelItemSkuDeleteRequest struct {
    model.Params
    // 宝贝ID
    _itemId   int64
    // 宝贝所属类目
    _catId   int64
    // 要删除的skuId列表
    _itemDeletedSkuIdList   []int64
}

// 初始化AlitripTuanHotelItemSkuDeleteRequest对象
func NewAlitripTuanHotelItemSkuDeleteRequest() *AlitripTuanHotelItemSkuDeleteRequest{
    return &AlitripTuanHotelItemSkuDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelItemSkuDeleteRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.item.sku.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelItemSkuDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 宝贝ID
func (r *AlitripTuanHotelItemSkuDeleteRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlitripTuanHotelItemSkuDeleteRequest) GetItemId() int64 {
    return r._itemId
}
// CatId Setter
// 宝贝所属类目
func (r *AlitripTuanHotelItemSkuDeleteRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlitripTuanHotelItemSkuDeleteRequest) GetCatId() int64 {
    return r._catId
}
// ItemDeletedSkuIdList Setter
// 要删除的skuId列表
func (r *AlitripTuanHotelItemSkuDeleteRequest) SetItemDeletedSkuIdList(_itemDeletedSkuIdList []int64) error {
    r._itemDeletedSkuIdList = _itemDeletedSkuIdList
    r.Set("item_deleted_sku_id_list", _itemDeletedSkuIdList)
    return nil
}

// ItemDeletedSkuIdList Getter
func (r AlitripTuanHotelItemSkuDeleteRequest) GetItemDeletedSkuIdList() []int64 {
    return r._itemDeletedSkuIdList
}
