package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店非标套餐商品日历库存宝贝SKU更新接口 APIRequest
alitrip.tuan.hotel.item.sku.calendar.update

商户对发布的日历库存类型的宝贝套餐价格库存信息进行更新，仅提供日历库存的宝贝SKU的更新功能，skuId须传递商品已存在的skuId，若想进行SKU新增操作，请选择使用alitrip.tuan.hotel.item.sku.update接口。提供增量更新SKU功能，对于日历库存若传递日期信息，参数中若包含某一日期的价格和库存，则对此日期的数据进行覆盖更新，若不传递则保留此日期的价格库存信息。
*/
type AlitripTuanHotelItemSkuCalendarUpdateRequest struct {
    model.Params

    // 宝贝ID
    itemId   int64 

    // 宝贝所属类目
    catId   int64 

    // 暂不支持此接口对SKU的部分属性进行更新，包括以下属性： 套餐名称、价格、原价、库存、间夜、商家编码、人数、使用次数等
    itemSkuList   []TopTuanItemSkuVO 

}

func NewAlitripTuanHotelItemSkuCalendarUpdateRequest() *AlitripTuanHotelItemSkuCalendarUpdateRequest{
    return &AlitripTuanHotelItemSkuCalendarUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTuanHotelItemSkuCalendarUpdateRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.item.sku.calendar.update"
}

func (r AlitripTuanHotelItemSkuCalendarUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTuanHotelItemSkuCalendarUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlitripTuanHotelItemSkuCalendarUpdateRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlitripTuanHotelItemSkuCalendarUpdateRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r AlitripTuanHotelItemSkuCalendarUpdateRequest) GetCatId() int64 {
    return r.catId
}

func (r *AlitripTuanHotelItemSkuCalendarUpdateRequest) SetItemSkuList(itemSkuList []TopTuanItemSkuVO) error {
    r.itemSkuList = itemSkuList
    r.Set("item_sku_list", itemSkuList)
    return nil
}

func (r AlitripTuanHotelItemSkuCalendarUpdateRequest) GetItemSkuList() []TopTuanItemSkuVO {
    return r.itemSkuList
}

