package uscesl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslIteminfoBatchInsertAPIRequest 按商家批量写入商品接口 API请求
// taobao.uscesl.iteminfo.batch.insert
//
// 【电子价签】支持按照商家-门店维度批量写入商品数据
type TaobaoUsceslIteminfoBatchInsertAPIRequest struct {
	model.Params
	// 商品信息集合，限制500
	_itemList []ItemChangeBo
	// 商家KEY
	_bizBrandKey string
	// 门店ID
	_storeId int64
}

// NewTaobaoUsceslIteminfoBatchInsertRequest 初始化TaobaoUsceslIteminfoBatchInsertAPIRequest对象
func NewTaobaoUsceslIteminfoBatchInsertRequest() *TaobaoUsceslIteminfoBatchInsertAPIRequest {
	return &TaobaoUsceslIteminfoBatchInsertAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsceslIteminfoBatchInsertAPIRequest) Reset() {
	r._itemList = r._itemList[:0]
	r._bizBrandKey = ""
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.iteminfo.batch.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemList is ItemList Setter
// 商品信息集合，限制500
func (r *TaobaoUsceslIteminfoBatchInsertAPIRequest) SetItemList(_itemList []ItemChangeBo) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetItemList() []ItemChangeBo {
	return r._itemList
}

// SetBizBrandKey is BizBrandKey Setter
// 商家KEY
func (r *TaobaoUsceslIteminfoBatchInsertAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoUsceslIteminfoBatchInsertAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoUsceslIteminfoBatchInsertAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsceslIteminfoBatchInsertRequest()
	},
}

// GetTaobaoUsceslIteminfoBatchInsertRequest 从 sync.Pool 获取 TaobaoUsceslIteminfoBatchInsertAPIRequest
func GetTaobaoUsceslIteminfoBatchInsertAPIRequest() *TaobaoUsceslIteminfoBatchInsertAPIRequest {
	return poolTaobaoUsceslIteminfoBatchInsertAPIRequest.Get().(*TaobaoUsceslIteminfoBatchInsertAPIRequest)
}

// ReleaseTaobaoUsceslIteminfoBatchInsertAPIRequest 将 TaobaoUsceslIteminfoBatchInsertAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsceslIteminfoBatchInsertAPIRequest(v *TaobaoUsceslIteminfoBatchInsertAPIRequest) {
	v.Reset()
	poolTaobaoUsceslIteminfoBatchInsertAPIRequest.Put(v)
}
