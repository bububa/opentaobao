package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品编辑提交schema信息 APIRequest
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

func NewAlibabaItemEditSubmitRequest() *AlibabaItemEditSubmitRequest{
    return &AlibabaItemEditSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItemEditSubmitRequest) GetApiMethodName() string {
    return "alibaba.item.edit.submit"
}

func (r AlibabaItemEditSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItemEditSubmitRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r AlibabaItemEditSubmitRequest) GetBizType() string {
    return r.bizType
}

func (r *AlibabaItemEditSubmitRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r AlibabaItemEditSubmitRequest) GetCatId() int64 {
    return r.catId
}

func (r *AlibabaItemEditSubmitRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

func (r AlibabaItemEditSubmitRequest) GetSpuId() int64 {
    return r.spuId
}

func (r *AlibabaItemEditSubmitRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaItemEditSubmitRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlibabaItemEditSubmitRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r AlibabaItemEditSubmitRequest) GetSchema() string {
    return r.schema
}

