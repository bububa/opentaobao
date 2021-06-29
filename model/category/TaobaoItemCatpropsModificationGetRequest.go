package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品类目属性变更 API请求
taobao.item.catprops.modification.get

查询商品类目属性变更信息
*/
type TaobaoItemCatpropsModificationGetRequest struct {
    model.Params
    // 类目Id（与商品Id二选一即可）
    categoryId   int64
    // 商品Id（与类目Id二选一即可。若同时传入商品Id和类目Id，则优先使用商品Id。若填写商品Id，则起始时间设为该商品最近修改时间）
    itemId   string
    // 起始请求时间（建议传入，默认为90天内）
    startTime   string
}

// 初始化TaobaoItemCatpropsModificationGetRequest对象
func NewTaobaoItemCatpropsModificationGetRequest() *TaobaoItemCatpropsModificationGetRequest{
    return &TaobaoItemCatpropsModificationGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemCatpropsModificationGetRequest) GetApiMethodName() string {
    return "taobao.item.catprops.modification.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemCatpropsModificationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// 类目Id（与商品Id二选一即可）
func (r *TaobaoItemCatpropsModificationGetRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r TaobaoItemCatpropsModificationGetRequest) GetCategoryId() int64 {
    return r.categoryId
}
// ItemId Setter
// 商品Id（与类目Id二选一即可。若同时传入商品Id和类目Id，则优先使用商品Id。若填写商品Id，则起始时间设为该商品最近修改时间）
func (r *TaobaoItemCatpropsModificationGetRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoItemCatpropsModificationGetRequest) GetItemId() string {
    return r.itemId
}
// StartTime Setter
// 起始请求时间（建议传入，默认为90天内）
func (r *TaobaoItemCatpropsModificationGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoItemCatpropsModificationGetRequest) GetStartTime() string {
    return r.startTime
}
