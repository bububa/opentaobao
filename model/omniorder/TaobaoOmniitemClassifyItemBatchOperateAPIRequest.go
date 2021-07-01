package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加/删除商品和分类的关联关系 API请求
taobao.omniitem.classify.item.batch.operate

批量添加/删除商品和分类的关联关系
*/
type TaobaoOmniitemClassifyItemBatchOperateAPIRequest struct {
    model.Params
    // 分类ID
    _classifyIds   []int64
    // 需要添加分类关联关系的商品ID
    _addItemIds   []int64
    // 需要删除分类关联关系的商品ID
    _deleteItemIds   []int64
    // 操作人信息（暂时不填）
    _operator   string
}

// 初始化TaobaoOmniitemClassifyItemBatchOperateAPIRequest对象
func NewTaobaoOmniitemClassifyItemBatchOperateRequest() *TaobaoOmniitemClassifyItemBatchOperateAPIRequest{
    return &TaobaoOmniitemClassifyItemBatchOperateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyItemBatchOperateAPIRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.item.batch.operate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyItemBatchOperateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClassifyIds Setter
// 分类ID
func (r *TaobaoOmniitemClassifyItemBatchOperateAPIRequest) SetClassifyIds(_classifyIds []int64) error {
    r._classifyIds = _classifyIds
    r.Set("classify_ids", _classifyIds)
    return nil
}

// ClassifyIds Getter
func (r TaobaoOmniitemClassifyItemBatchOperateAPIRequest) GetClassifyIds() []int64 {
    return r._classifyIds
}
// AddItemIds Setter
// 需要添加分类关联关系的商品ID
func (r *TaobaoOmniitemClassifyItemBatchOperateAPIRequest) SetAddItemIds(_addItemIds []int64) error {
    r._addItemIds = _addItemIds
    r.Set("add_item_ids", _addItemIds)
    return nil
}

// AddItemIds Getter
func (r TaobaoOmniitemClassifyItemBatchOperateAPIRequest) GetAddItemIds() []int64 {
    return r._addItemIds
}
// DeleteItemIds Setter
// 需要删除分类关联关系的商品ID
func (r *TaobaoOmniitemClassifyItemBatchOperateAPIRequest) SetDeleteItemIds(_deleteItemIds []int64) error {
    r._deleteItemIds = _deleteItemIds
    r.Set("delete_item_ids", _deleteItemIds)
    return nil
}

// DeleteItemIds Getter
func (r TaobaoOmniitemClassifyItemBatchOperateAPIRequest) GetDeleteItemIds() []int64 {
    return r._deleteItemIds
}
// Operator Setter
// 操作人信息（暂时不填）
func (r *TaobaoOmniitemClassifyItemBatchOperateAPIRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoOmniitemClassifyItemBatchOperateAPIRequest) GetOperator() string {
    return r._operator
}
