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
    _catId   int64
    // 产品ID，若不需要修改关联的产品信息，则不需要填写
    _spuId   int64
    // 商品ID
    _itemId   int64
    // 编辑后的schema信息(增量更新，只填写需要更新的字段)
    _schema   string
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
func (r *AlibabaItemEditFastupdateRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlibabaItemEditFastupdateRequest) GetCatId() int64 {
    return r._catId
}
// SpuId Setter
// 产品ID，若不需要修改关联的产品信息，则不需要填写
func (r *AlibabaItemEditFastupdateRequest) SetSpuId(_spuId int64) error {
    r._spuId = _spuId
    r.Set("spu_id", _spuId)
    return nil
}

// SpuId Getter
func (r AlibabaItemEditFastupdateRequest) GetSpuId() int64 {
    return r._spuId
}
// ItemId Setter
// 商品ID
func (r *AlibabaItemEditFastupdateRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlibabaItemEditFastupdateRequest) GetItemId() int64 {
    return r._itemId
}
// Schema Setter
// 编辑后的schema信息(增量更新，只填写需要更新的字段)
func (r *AlibabaItemEditFastupdateRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r AlibabaItemEditFastupdateRequest) GetSchema() string {
    return r._schema
}
