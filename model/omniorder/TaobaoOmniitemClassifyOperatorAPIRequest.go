package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyOperatorAPIRequest
添加/修改分类 API请求
taobao.omniitem.classify.operator

添加/修改分类 */
type TaobaoOmniitemClassifyOperatorAPIRequest struct {
	model.Params
	// 分类信息
	_category *OmniItemCategoryDetailDto
	// 需要添加的关联关系的商品
	_addItemIds []int64
	// 需要修改的关联关系的商品
	_removeItemIds []int64
	// 操作人信息（暂时不填）
	_operator string
}

// NewTaobaoOmniitemClassifyOperatorRequest 初始化TaobaoOmniitemClassifyOperatorAPIRequest对象
func NewTaobaoOmniitemClassifyOperatorRequest() *TaobaoOmniitemClassifyOperatorAPIRequest {
	return &TaobaoOmniitemClassifyOperatorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyOperatorAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.classify.operator"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyOperatorAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Category Setter
// 分类信息
func (r *TaobaoOmniitemClassifyOperatorAPIRequest) SetCategory(_category *OmniItemCategoryDetailDto) error {
	r._category = _category
	r.Set("category", _category)
	return nil
}

// Get Category Getter
func (r TaobaoOmniitemClassifyOperatorAPIRequest) GetCategory() *OmniItemCategoryDetailDto {
	return r._category
}

// Set is AddItemIds Setter
// 需要添加的关联关系的商品
func (r *TaobaoOmniitemClassifyOperatorAPIRequest) SetAddItemIds(_addItemIds []int64) error {
	r._addItemIds = _addItemIds
	r.Set("add_item_ids", _addItemIds)
	return nil
}

// Get AddItemIds Getter
func (r TaobaoOmniitemClassifyOperatorAPIRequest) GetAddItemIds() []int64 {
	return r._addItemIds
}

// Set is RemoveItemIds Setter
// 需要修改的关联关系的商品
func (r *TaobaoOmniitemClassifyOperatorAPIRequest) SetRemoveItemIds(_removeItemIds []int64) error {
	r._removeItemIds = _removeItemIds
	r.Set("remove_item_ids", _removeItemIds)
	return nil
}

// Get RemoveItemIds Getter
func (r TaobaoOmniitemClassifyOperatorAPIRequest) GetRemoveItemIds() []int64 {
	return r._removeItemIds
}

// Set is Operator Setter
// 操作人信息（暂时不填）
func (r *TaobaoOmniitemClassifyOperatorAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// Get Operator Getter
func (r TaobaoOmniitemClassifyOperatorAPIRequest) GetOperator() string {
	return r._operator
}
