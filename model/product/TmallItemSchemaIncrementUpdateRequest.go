package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫根据规则增量更新商品 APIRequest
tmall.item.schema.increment.update

增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名）
*/
type TmallItemSchemaIncrementUpdateRequest struct {
    model.Params

    // 需要编辑的商品ID
    itemId   int64 

    // 根据tmall.item.increment.update.schema.get生成的商品增量编辑规则入参数据。需要更新的字段，一定要在入参的XML重点update_fields字段中明确指明
    xmlData   string 

}

func NewTmallItemSchemaIncrementUpdateRequest() *TmallItemSchemaIncrementUpdateRequest{
    return &TmallItemSchemaIncrementUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemSchemaIncrementUpdateRequest) GetApiMethodName() string {
    return "tmall.item.schema.increment.update"
}

func (r TmallItemSchemaIncrementUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemSchemaIncrementUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemSchemaIncrementUpdateRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TmallItemSchemaIncrementUpdateRequest) SetXmlData(xmlData string) error {
    r.xmlData = xmlData
    r.Set("xml_data", xmlData)
    return nil
}

func (r TmallItemSchemaIncrementUpdateRequest) GetXmlData() string {
    return r.xmlData
}

