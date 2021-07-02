package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtItemGetAPIRequest 家装新零售商品信息查询 API请求
// tmall.nrt.item.get
//
// 查询新零售商品信息
type TmallNrtItemGetAPIRequest struct {
	model.Params
	// 城市站id
	_boothId int64
	// 商品id
	_itemId int64
}

// NewTmallNrtItemGetRequest 初始化TmallNrtItemGetAPIRequest对象
func NewTmallNrtItemGetRequest() *TmallNrtItemGetAPIRequest {
	return &TmallNrtItemGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtItemGetAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtItemGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBoothId is BoothId Setter
// 城市站id
func (r *TmallNrtItemGetAPIRequest) SetBoothId(_boothId int64) error {
	r._boothId = _boothId
	r.Set("booth_id", _boothId)
	return nil
}

// GetBoothId BoothId Getter
func (r TmallNrtItemGetAPIRequest) GetBoothId() int64 {
	return r._boothId
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallNrtItemGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallNrtItemGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
