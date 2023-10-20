package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctionbeikeitemsyncAPIRequest 贝壳商品同步接口 API请求
// taobao.auction.beike.item.sync
//
// 贝壳商品同步接口
type TaobaoauctionbeikeitemsyncAPIRequest struct {
	model.Params
	// 日期
	_ds int64
	// 无
	_beikeItemDo *BeikeItemDo
}

// NewTaobaoauctionbeikeitemsyncRequest 初始化TaobaoauctionbeikeitemsyncAPIRequest对象
func NewTaobaoauctionbeikeitemsyncRequest() *TaobaoauctionbeikeitemsyncAPIRequest {
	return &TaobaoauctionbeikeitemsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoauctionbeikeitemsyncAPIRequest) GetApiMethodName() string {
	return "taobao.auction.beike.item.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoauctionbeikeitemsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoauctionbeikeitemsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDs is Ds Setter
// 日期
func (r *TaobaoauctionbeikeitemsyncAPIRequest) SetDs(_ds int64) error {
	r._ds = _ds
	r.Set("ds", _ds)
	return nil
}

// GetDs Ds Getter
func (r TaobaoauctionbeikeitemsyncAPIRequest) GetDs() int64 {
	return r._ds
}

// SetBeikeItemDo is BeikeItemDo Setter
// 无
func (r *TaobaoauctionbeikeitemsyncAPIRequest) SetBeikeItemDo(_beikeItemDo *BeikeItemDo) error {
	r._beikeItemDo = _beikeItemDo
	r.Set("beike_item_do", _beikeItemDo)
	return nil
}

// GetBeikeItemDo BeikeItemDo Getter
func (r TaobaoauctionbeikeitemsyncAPIRequest) GetBeikeItemDo() *BeikeItemDo {
	return r._beikeItemDo
}
