package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallItemsQueryAPIRequest 批量获取商品列表 API请求
// taobao.openmall.items.query
//
// 批量获取对联盟开放的商品列表。
type TaobaoOpenmallItemsQueryAPIRequest struct {
	model.Params
	// 已废弃，请勿使用
	_itemIds string
	// 当不输入渠道商时，展示全网公有商品池；当输入渠道商的淘宝Nick时，展示该渠道私有供给商品列表
	_distributor string
	// 第几页，默认：1
	_pageNo int64
	// 页大小，默认20，1~100
	_pageSize int64
}

// NewTaobaoOpenmallItemsQueryRequest 初始化TaobaoOpenmallItemsQueryAPIRequest对象
func NewTaobaoOpenmallItemsQueryRequest() *TaobaoOpenmallItemsQueryAPIRequest {
	return &TaobaoOpenmallItemsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallItemsQueryAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.items.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallItemsQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemIds is ItemIds Setter
// 已废弃，请勿使用
func (r *TaobaoOpenmallItemsQueryAPIRequest) SetItemIds(_itemIds string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoOpenmallItemsQueryAPIRequest) GetItemIds() string {
	return r._itemIds
}

// SetDistributor is Distributor Setter
// 当不输入渠道商时，展示全网公有商品池；当输入渠道商的淘宝Nick时，展示该渠道私有供给商品列表
func (r *TaobaoOpenmallItemsQueryAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallItemsQueryAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetPageNo is PageNo Setter
// 第几页，默认：1
func (r *TaobaoOpenmallItemsQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoOpenmallItemsQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoOpenmallItemsQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoOpenmallItemsQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
