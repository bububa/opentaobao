package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemShelveAPIRequest 【API3.0】度假线路商品上下架接口 API请求
// taobao.alitrip.travel.item.shelve
//
// 旅行度假新商品发布接口 第三版：度假商品上下架接口
// 注意：定时上下架功能，目前只支持接送、租车类目。
type TaobaoAlitripTravelItemShelveAPIRequest struct {
	model.Params
	// 指定定时上架时间，格式：yyyy-MM-dd HH:mm:ss。若不设置该值且item_status为1，则表示立即上架。
	_onlineTime string
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 1-上架 0-下架
	_itemStatus int64
}

// NewTaobaoAlitripTravelItemShelveRequest 初始化TaobaoAlitripTravelItemShelveAPIRequest对象
func NewTaobaoAlitripTravelItemShelveRequest() *TaobaoAlitripTravelItemShelveAPIRequest {
	return &TaobaoAlitripTravelItemShelveAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelItemShelveAPIRequest) Reset() {
	r._onlineTime = ""
	r._outProductId = ""
	r._itemId = 0
	r._itemStatus = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemShelveAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.shelve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemShelveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelItemShelveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOnlineTime is OnlineTime Setter
// 指定定时上架时间，格式：yyyy-MM-dd HH:mm:ss。若不设置该值且item_status为1，则表示立即上架。
func (r *TaobaoAlitripTravelItemShelveAPIRequest) SetOnlineTime(_onlineTime string) error {
	r._onlineTime = _onlineTime
	r.Set("online_time", _onlineTime)
	return nil
}

// GetOnlineTime OnlineTime Getter
func (r TaobaoAlitripTravelItemShelveAPIRequest) GetOnlineTime() string {
	return r._onlineTime
}

// SetOutProductId is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemShelveAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r TaobaoAlitripTravelItemShelveAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemShelveAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoAlitripTravelItemShelveAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemStatus is ItemStatus Setter
// 1-上架 0-下架
func (r *TaobaoAlitripTravelItemShelveAPIRequest) SetItemStatus(_itemStatus int64) error {
	r._itemStatus = _itemStatus
	r.Set("item_status", _itemStatus)
	return nil
}

// GetItemStatus ItemStatus Getter
func (r TaobaoAlitripTravelItemShelveAPIRequest) GetItemStatus() int64 {
	return r._itemStatus
}

var poolTaobaoAlitripTravelItemShelveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelItemShelveRequest()
	},
}

// GetTaobaoAlitripTravelItemShelveRequest 从 sync.Pool 获取 TaobaoAlitripTravelItemShelveAPIRequest
func GetTaobaoAlitripTravelItemShelveAPIRequest() *TaobaoAlitripTravelItemShelveAPIRequest {
	return poolTaobaoAlitripTravelItemShelveAPIRequest.Get().(*TaobaoAlitripTravelItemShelveAPIRequest)
}

// ReleaseTaobaoAlitripTravelItemShelveAPIRequest 将 TaobaoAlitripTravelItemShelveAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelItemShelveAPIRequest(v *TaobaoAlitripTravelItemShelveAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelItemShelveAPIRequest.Put(v)
}
