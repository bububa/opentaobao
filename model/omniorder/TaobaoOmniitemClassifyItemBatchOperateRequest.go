package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加/删除商品和分类的关联关系 APIRequest
taobao.omniitem.classify.item.batch.operate

批量添加/删除商品和分类的关联关系
*/
type TaobaoOmniitemClassifyItemBatchOperateRequest struct {
    model.Params

    // 分类ID
    classifyIds   []int64 

    // 需要添加分类关联关系的商品ID
    addItemIds   []int64 

    // 需要删除分类关联关系的商品ID
    deleteItemIds   []int64 

    // 操作人信息（暂时不填）
    operator   string 

}

func NewTaobaoOmniitemClassifyItemBatchOperateRequest() *TaobaoOmniitemClassifyItemBatchOperateRequest{
    return &TaobaoOmniitemClassifyItemBatchOperateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemClassifyItemBatchOperateRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.item.batch.operate"
}

func (r TaobaoOmniitemClassifyItemBatchOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemClassifyItemBatchOperateRequest) SetClassifyIds(classifyIds []int64) error {
    r.classifyIds = classifyIds
    r.Set("classify_ids", classifyIds)
    return nil
}

func (r TaobaoOmniitemClassifyItemBatchOperateRequest) GetClassifyIds() []int64 {
    return r.classifyIds
}

func (r *TaobaoOmniitemClassifyItemBatchOperateRequest) SetAddItemIds(addItemIds []int64) error {
    r.addItemIds = addItemIds
    r.Set("add_item_ids", addItemIds)
    return nil
}

func (r TaobaoOmniitemClassifyItemBatchOperateRequest) GetAddItemIds() []int64 {
    return r.addItemIds
}

func (r *TaobaoOmniitemClassifyItemBatchOperateRequest) SetDeleteItemIds(deleteItemIds []int64) error {
    r.deleteItemIds = deleteItemIds
    r.Set("delete_item_ids", deleteItemIds)
    return nil
}

func (r TaobaoOmniitemClassifyItemBatchOperateRequest) GetDeleteItemIds() []int64 {
    return r.deleteItemIds
}

func (r *TaobaoOmniitemClassifyItemBatchOperateRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoOmniitemClassifyItemBatchOperateRequest) GetOperator() string {
    return r.operator
}

