package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品编辑获取schema信息 API请求
alibaba.item.edit.schema.get

商品编辑时，获取商品规则信息
*/
type AlibabaItemEditSchemaGetRequest struct {
    model.Params
    // 业务扩展参数，需与平台约定好
    bizType   string
    // 商品ID
    itemId   int64
    // 制定返回schema中field字段列表，可用于裁剪返回的schema信息。不填则为全部field
    fields   []string
}

// 初始化AlibabaItemEditSchemaGetRequest对象
func NewAlibabaItemEditSchemaGetRequest() *AlibabaItemEditSchemaGetRequest{
    return &AlibabaItemEditSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemEditSchemaGetRequest) GetApiMethodName() string {
    return "alibaba.item.edit.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemEditSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务扩展参数，需与平台约定好
func (r *AlibabaItemEditSchemaGetRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r AlibabaItemEditSchemaGetRequest) GetBizType() string {
    return r.bizType
}
// ItemId Setter
// 商品ID
func (r *AlibabaItemEditSchemaGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaItemEditSchemaGetRequest) GetItemId() int64 {
    return r.itemId
}
// Fields Setter
// 制定返回schema中field字段列表，可用于裁剪返回的schema信息。不填则为全部field
func (r *AlibabaItemEditSchemaGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r AlibabaItemEditSchemaGetRequest) GetFields() []string {
    return r.fields
}
