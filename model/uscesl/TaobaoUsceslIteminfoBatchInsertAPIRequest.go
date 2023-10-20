package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouscesliteminfobatchinsertAPIRequest 按商家批量写入商品接口 API请求
// taobao.uscesl.iteminfo.batch.insert
//
// 【电子价签】支持按照商家-门店维度批量写入商品数据
type TaobaouscesliteminfobatchinsertAPIRequest struct {
	model.Params
	// 商品信息集合，限制500
	_itemList []ItemChangeBo
	// 商家KEY
	_bizBrandKey string
	// 门店ID
	_storeId int64
}

// NewTaobaouscesliteminfobatchinsertRequest 初始化TaobaouscesliteminfobatchinsertAPIRequest对象
func NewTaobaouscesliteminfobatchinsertRequest() *TaobaouscesliteminfobatchinsertAPIRequest {
	return &TaobaouscesliteminfobatchinsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouscesliteminfobatchinsertAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.iteminfo.batch.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouscesliteminfobatchinsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouscesliteminfobatchinsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemList is ItemList Setter
// 商品信息集合，限制500
func (r *TaobaouscesliteminfobatchinsertAPIRequest) SetItemList(_itemList []ItemChangeBo) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r TaobaouscesliteminfobatchinsertAPIRequest) GetItemList() []ItemChangeBo {
	return r._itemList
}

// SetBizBrandKey is BizBrandKey Setter
// 商家KEY
func (r *TaobaouscesliteminfobatchinsertAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaouscesliteminfobatchinsertAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaouscesliteminfobatchinsertAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaouscesliteminfobatchinsertAPIRequest) GetStoreId() int64 {
	return r._storeId
}
