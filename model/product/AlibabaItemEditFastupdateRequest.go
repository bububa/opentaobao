package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品编辑增量更新 APIRequest
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

func NewAlibabaItemEditFastupdateRequest() *AlibabaItemEditFastupdateRequest{
    return &AlibabaItemEditFastupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItemEditFastupdateRequest) GetApiMethodName() string {
    return "alibaba.item.edit.fastupdate"
}

func (r AlibabaItemEditFastupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItemEditFastupdateRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r AlibabaItemEditFastupdateRequest) GetCatId() int64 {
    return r.catId
}

func (r *AlibabaItemEditFastupdateRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

func (r AlibabaItemEditFastupdateRequest) GetSpuId() int64 {
    return r.spuId
}

func (r *AlibabaItemEditFastupdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaItemEditFastupdateRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlibabaItemEditFastupdateRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r AlibabaItemEditFastupdateRequest) GetSchema() string {
    return r.schema
}

