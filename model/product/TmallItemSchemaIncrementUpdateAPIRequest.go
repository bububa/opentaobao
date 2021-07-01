package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫根据规则增量更新商品 API请求
tmall.item.schema.increment.update

增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名）
*/
type TmallItemSchemaIncrementUpdateAPIRequest struct {
    model.Params
    // 需要编辑的商品ID
    _itemId   int64
    // 根据tmall.item.increment.update.schema.get生成的商品增量编辑规则入参数据。需要更新的字段，一定要在入参的XML重点update_fields字段中明确指明
    _xmlData   string
}

// 初始化TmallItemSchemaIncrementUpdateAPIRequest对象
func NewTmallItemSchemaIncrementUpdateRequest() *TmallItemSchemaIncrementUpdateAPIRequest{
    return &TmallItemSchemaIncrementUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSchemaIncrementUpdateAPIRequest) GetApiMethodName() string {
    return "tmall.item.schema.increment.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSchemaIncrementUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 需要编辑的商品ID
func (r *TmallItemSchemaIncrementUpdateAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallItemSchemaIncrementUpdateAPIRequest) GetItemId() int64 {
    return r._itemId
}
// XmlData Setter
// 根据tmall.item.increment.update.schema.get生成的商品增量编辑规则入参数据。需要更新的字段，一定要在入参的XML重点update_fields字段中明确指明
func (r *TmallItemSchemaIncrementUpdateAPIRequest) SetXmlData(_xmlData string) error {
    r._xmlData = _xmlData
    r.Set("xml_data", _xmlData)
    return nil
}

// XmlData Getter
func (r TmallItemSchemaIncrementUpdateAPIRequest) GetXmlData() string {
    return r._xmlData
}
