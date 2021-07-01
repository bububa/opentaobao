package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyStoreBatchOperateAPIRequest
批量添加/删除门店和分类的关联关系 API请求
taobao.omniitem.classify.store.batch.operate

批量添加/删除门店和分类的关联关系 */
type TaobaoOmniitemClassifyStoreBatchOperateAPIRequest struct {
	model.Params
	// 商家门店ID
	_storeIds []int64
	// 需要添加的分类ID
	_addCategoryIds []int64
	// 需要删除的分类ID
	_removeCategoryIds []int64
	// 操作信息（暂时不填）
	_operator string
}

// NewTaobaoOmniitemClassifyStoreBatchOperateRequest 初始化TaobaoOmniitemClassifyStoreBatchOperateAPIRequest对象
func NewTaobaoOmniitemClassifyStoreBatchOperateRequest() *TaobaoOmniitemClassifyStoreBatchOperateAPIRequest {
	return &TaobaoOmniitemClassifyStoreBatchOperateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyStoreBatchOperateAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.classify.store.batch.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyStoreBatchOperateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreIds Setter
// 商家门店ID
func (r *TaobaoOmniitemClassifyStoreBatchOperateAPIRequest) SetStoreIds(_storeIds []int64) error {
	r._storeIds = _storeIds
	r.Set("store_ids", _storeIds)
	return nil
}

// Get StoreIds Getter
func (r TaobaoOmniitemClassifyStoreBatchOperateAPIRequest) GetStoreIds() []int64 {
	return r._storeIds
}

// Set is AddCategoryIds Setter
// 需要添加的分类ID
func (r *TaobaoOmniitemClassifyStoreBatchOperateAPIRequest) SetAddCategoryIds(_addCategoryIds []int64) error {
	r._addCategoryIds = _addCategoryIds
	r.Set("add_category_ids", _addCategoryIds)
	return nil
}

// Get AddCategoryIds Getter
func (r TaobaoOmniitemClassifyStoreBatchOperateAPIRequest) GetAddCategoryIds() []int64 {
	return r._addCategoryIds
}

// Set is RemoveCategoryIds Setter
// 需要删除的分类ID
func (r *TaobaoOmniitemClassifyStoreBatchOperateAPIRequest) SetRemoveCategoryIds(_removeCategoryIds []int64) error {
	r._removeCategoryIds = _removeCategoryIds
	r.Set("remove_category_ids", _removeCategoryIds)
	return nil
}

// Get RemoveCategoryIds Getter
func (r TaobaoOmniitemClassifyStoreBatchOperateAPIRequest) GetRemoveCategoryIds() []int64 {
	return r._removeCategoryIds
}

// Set is Operator Setter
// 操作信息（暂时不填）
func (r *TaobaoOmniitemClassifyStoreBatchOperateAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// Get Operator Getter
func (r TaobaoOmniitemClassifyStoreBatchOperateAPIRequest) GetOperator() string {
	return r._operator
}
