package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemCatpropsModificationGetAPIRequest 查询商品类目属性变更 API请求
// taobao.item.catprops.modification.get
//
// 查询商品类目属性变更信息
type TaobaoItemCatpropsModificationGetAPIRequest struct {
	model.Params
	// 类目Id（与商品Id二选一即可）
	_categoryId int64
	// 商品Id（与类目Id二选一即可。若同时传入商品Id和类目Id，则优先使用商品Id。若填写商品Id，则起始时间设为该商品最近修改时间）
	_itemId string
	// 起始请求时间（建议传入，默认为90天内）
	_startTime string
}

// NewTaobaoItemCatpropsModificationGetRequest 初始化TaobaoItemCatpropsModificationGetAPIRequest对象
func NewTaobaoItemCatpropsModificationGetRequest() *TaobaoItemCatpropsModificationGetAPIRequest {
	return &TaobaoItemCatpropsModificationGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemCatpropsModificationGetAPIRequest) GetApiMethodName() string {
	return "taobao.item.catprops.modification.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemCatpropsModificationGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CategoryId Setter
// 类目Id（与商品Id二选一即可）
func (r *TaobaoItemCatpropsModificationGetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// Get CategoryId Getter
func (r TaobaoItemCatpropsModificationGetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// Set is ItemId Setter
// 商品Id（与类目Id二选一即可。若同时传入商品Id和类目Id，则优先使用商品Id。若填写商品Id，则起始时间设为该商品最近修改时间）
func (r *TaobaoItemCatpropsModificationGetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoItemCatpropsModificationGetAPIRequest) GetItemId() string {
	return r._itemId
}

// Set is StartTime Setter
// 起始请求时间（建议传入，默认为90天内）
func (r *TaobaoItemCatpropsModificationGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoItemCatpropsModificationGetAPIRequest) GetStartTime() string {
	return r._startTime
}
