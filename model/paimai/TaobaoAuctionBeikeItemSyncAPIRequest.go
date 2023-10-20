package paimai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionBeikeItemSyncAPIRequest 贝壳商品同步接口 API请求
// taobao.auction.beike.item.sync
//
// 贝壳商品同步接口
type TaobaoAuctionBeikeItemSyncAPIRequest struct {
	model.Params
	// 日期
	_ds int64
	// 无
	_beikeItemDo *BeikeItemDo
}

// NewTaobaoAuctionBeikeItemSyncRequest 初始化TaobaoAuctionBeikeItemSyncAPIRequest对象
func NewTaobaoAuctionBeikeItemSyncRequest() *TaobaoAuctionBeikeItemSyncAPIRequest {
	return &TaobaoAuctionBeikeItemSyncAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAuctionBeikeItemSyncAPIRequest) Reset() {
	r._ds = 0
	r._beikeItemDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionBeikeItemSyncAPIRequest) GetApiMethodName() string {
	return "taobao.auction.beike.item.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionBeikeItemSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAuctionBeikeItemSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDs is Ds Setter
// 日期
func (r *TaobaoAuctionBeikeItemSyncAPIRequest) SetDs(_ds int64) error {
	r._ds = _ds
	r.Set("ds", _ds)
	return nil
}

// GetDs Ds Getter
func (r TaobaoAuctionBeikeItemSyncAPIRequest) GetDs() int64 {
	return r._ds
}

// SetBeikeItemDo is BeikeItemDo Setter
// 无
func (r *TaobaoAuctionBeikeItemSyncAPIRequest) SetBeikeItemDo(_beikeItemDo *BeikeItemDo) error {
	r._beikeItemDo = _beikeItemDo
	r.Set("beike_item_do", _beikeItemDo)
	return nil
}

// GetBeikeItemDo BeikeItemDo Getter
func (r TaobaoAuctionBeikeItemSyncAPIRequest) GetBeikeItemDo() *BeikeItemDo {
	return r._beikeItemDo
}

var poolTaobaoAuctionBeikeItemSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAuctionBeikeItemSyncRequest()
	},
}

// GetTaobaoAuctionBeikeItemSyncRequest 从 sync.Pool 获取 TaobaoAuctionBeikeItemSyncAPIRequest
func GetTaobaoAuctionBeikeItemSyncAPIRequest() *TaobaoAuctionBeikeItemSyncAPIRequest {
	return poolTaobaoAuctionBeikeItemSyncAPIRequest.Get().(*TaobaoAuctionBeikeItemSyncAPIRequest)
}

// ReleaseTaobaoAuctionBeikeItemSyncAPIRequest 将 TaobaoAuctionBeikeItemSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoAuctionBeikeItemSyncAPIRequest(v *TaobaoAuctionBeikeItemSyncAPIRequest) {
	v.Reset()
	poolTaobaoAuctionBeikeItemSyncAPIRequest.Put(v)
}
