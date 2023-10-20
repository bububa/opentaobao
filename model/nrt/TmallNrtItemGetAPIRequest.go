package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtitemgetAPIRequest 家装新零售商品信息查询 API请求
// tmall.nrt.item.get
//
// 查询新零售商品信息
type TmallnrtitemgetAPIRequest struct {
	model.Params
	// 城市站id
	_boothId int64
	// 商品id
	_itemId int64
}

// NewTmallnrtitemgetRequest 初始化TmallnrtitemgetAPIRequest对象
func NewTmallnrtitemgetRequest() *TmallnrtitemgetAPIRequest {
	return &TmallnrtitemgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtitemgetAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtitemgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtitemgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBoothId is BoothId Setter
// 城市站id
func (r *TmallnrtitemgetAPIRequest) SetBoothId(_boothId int64) error {
	r._boothId = _boothId
	r.Set("booth_id", _boothId)
	return nil
}

// GetBoothId BoothId Getter
func (r TmallnrtitemgetAPIRequest) GetBoothId() int64 {
	return r._boothId
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallnrtitemgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallnrtitemgetAPIRequest) GetItemId() int64 {
	return r._itemId
}
