package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品编辑提交schema信息 API请求
alibaba.item.edit.submit

商品编辑提交schema信息
*/
type AlibabaItemEditSubmitRequest struct {
    model.Params
    // 业务扩展参数，需与平台约定好
    bizType   string
    // 商品类目ID。若不需要修改商品类目，则不用填写
    catId   int64
    // 产品ID，若不需要修改关联的产品信息，则不需要填写
    spuId   int64
    // 商品ID
    itemId   int64
    // 编辑后的schema信息，通过alibaba.item.edit.schema.get获取
    schema   string
}

// 初始化AlibabaItemEditSubmitRequest对象
func NewAlibabaItemEditSubmitRequest() *AlibabaItemEditSubmitRequest{
    return &AlibabaItemEditSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemEditSubmitRequest) GetApiMethodName() string {
    return "alibaba.item.edit.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemEditSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务扩展参数，需与平台约定好
func (r *AlibabaItemEditSubmitRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r AlibabaItemEditSubmitRequest) GetBizType() string {
    return r.bizType
}
// CatId Setter
// 商品类目ID。若不需要修改商品类目，则不用填写
func (r *AlibabaItemEditSubmitRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r AlibabaItemEditSubmitRequest) GetCatId() int64 {
    return r.catId
}
// SpuId Setter
// 产品ID，若不需要修改关联的产品信息，则不需要填写
func (r *AlibabaItemEditSubmitRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

// SpuId Getter
func (r AlibabaItemEditSubmitRequest) GetSpuId() int64 {
    return r.spuId
}
// ItemId Setter
// 商品ID
func (r *AlibabaItemEditSubmitRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaItemEditSubmitRequest) GetItemId() int64 {
    return r.itemId
}
// Schema Setter
// 编辑后的schema信息，通过alibaba.item.edit.schema.get获取
func (r *AlibabaItemEditSubmitRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r AlibabaItemEditSubmitRequest) GetSchema() string {
    return r.schema
}
