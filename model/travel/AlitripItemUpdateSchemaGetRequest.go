package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取编辑商品的schema模板 API请求
alitrip.item.update.schema.get

获取编辑商品的schema模板。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemUpdateSchemaGetRequest struct {
    model.Params
    // 商品id
    _itemId   int64
    // 需要获取的编辑schema，不填默认返回全部
    _updateFieldIds   []string
}

// 初始化AlitripItemUpdateSchemaGetRequest对象
func NewAlitripItemUpdateSchemaGetRequest() *AlitripItemUpdateSchemaGetRequest{
    return &AlitripItemUpdateSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripItemUpdateSchemaGetRequest) GetApiMethodName() string {
    return "alitrip.item.update.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripItemUpdateSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *AlitripItemUpdateSchemaGetRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlitripItemUpdateSchemaGetRequest) GetItemId() int64 {
    return r._itemId
}
// UpdateFieldIds Setter
// 需要获取的编辑schema，不填默认返回全部
func (r *AlitripItemUpdateSchemaGetRequest) SetUpdateFieldIds(_updateFieldIds []string) error {
    r._updateFieldIds = _updateFieldIds
    r.Set("update_field_ids", _updateFieldIds)
    return nil
}

// UpdateFieldIds Getter
func (r AlitripItemUpdateSchemaGetRequest) GetUpdateFieldIds() []string {
    return r._updateFieldIds
}
