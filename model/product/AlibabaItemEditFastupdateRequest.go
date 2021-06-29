package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品编辑增量更新 API请求
alibaba.item.edit.fastupdate

商品编辑增量更新;
<br/>该接口编辑sku，只能更新价格、库存等信息，不能新增sku;
<br/>新增sku用全量接口alibaba.item.edit.submit，先设置销售属性;
*/
type AlibabaItemEditFastupdateRequest struct {
    model.Params
    // 商品类目ID。若不需要修改商品类目，则不用填写
    catId   int64
    // 产品ID，若不需要修改关联的产品信息，则不需要填写
    spuId   int64
    // 商品ID
    itemId   int64
    // 编辑后的schema信息(增量更新，只填写需要更新的字段)
    schema   string
}

// 初始化AlibabaItemEditFastupdateRequest对象
func NewAlibabaItemEditFastupdateRequest() *AlibabaItemEditFastupdateRequest{
    return &AlibabaItemEditFastupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemEditFastupdateRequest) GetApiMethodName() string {
    return "alibaba.item.edit.fastupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemEditFastupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 商品类目ID。若不需要修改商品类目，则不用填写
func (r *AlibabaItemEditFastupdateRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r AlibabaItemEditFastupdateRequest) GetCatId() int64 {
    return r.catId
}
// SpuId Setter
// 产品ID，若不需要修改关联的产品信息，则不需要填写
func (r *AlibabaItemEditFastupdateRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

// SpuId Getter
func (r AlibabaItemEditFastupdateRequest) GetSpuId() int64 {
    return r.spuId
}
// ItemId Setter
// 商品ID
func (r *AlibabaItemEditFastupdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaItemEditFastupdateRequest) GetItemId() int64 {
    return r.itemId
}
// Schema Setter
// 编辑后的schema信息(增量更新，只填写需要更新的字段)
func (r *AlibabaItemEditFastupdateRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r AlibabaItemEditFastupdateRequest) GetSchema() string {
    return r.schema
}
