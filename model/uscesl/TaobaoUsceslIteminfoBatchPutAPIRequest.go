package uscesl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslIteminfoBatchPutAPIRequest 批量写入商品信息接口 API请求
// taobao.uscesl.iteminfo.batch.put
//
// 电子架签批量写入商品数据，用于电子价签展示
type TaobaoUsceslIteminfoBatchPutAPIRequest struct {
	model.Params
	// 商品变更信息集合
	_itemChangeBOList []ItemChangeBo
	// 门店ID
	_shopId int64
}

// NewTaobaoUsceslIteminfoBatchPutRequest 初始化TaobaoUsceslIteminfoBatchPutAPIRequest对象
func NewTaobaoUsceslIteminfoBatchPutRequest() *TaobaoUsceslIteminfoBatchPutAPIRequest {
	return &TaobaoUsceslIteminfoBatchPutAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsceslIteminfoBatchPutAPIRequest) Reset() {
	r._itemChangeBOList = r._itemChangeBOList[:0]
	r._shopId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslIteminfoBatchPutAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.iteminfo.batch.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslIteminfoBatchPutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsceslIteminfoBatchPutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemChangeBOList is ItemChangeBOList Setter
// 商品变更信息集合
func (r *TaobaoUsceslIteminfoBatchPutAPIRequest) SetItemChangeBOList(_itemChangeBOList []ItemChangeBo) error {
	r._itemChangeBOList = _itemChangeBOList
	r.Set("item_change_b_o_list", _itemChangeBOList)
	return nil
}

// GetItemChangeBOList ItemChangeBOList Getter
func (r TaobaoUsceslIteminfoBatchPutAPIRequest) GetItemChangeBOList() []ItemChangeBo {
	return r._itemChangeBOList
}

// SetShopId is ShopId Setter
// 门店ID
func (r *TaobaoUsceslIteminfoBatchPutAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaoUsceslIteminfoBatchPutAPIRequest) GetShopId() int64 {
	return r._shopId
}

var poolTaobaoUsceslIteminfoBatchPutAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsceslIteminfoBatchPutRequest()
	},
}

// GetTaobaoUsceslIteminfoBatchPutRequest 从 sync.Pool 获取 TaobaoUsceslIteminfoBatchPutAPIRequest
func GetTaobaoUsceslIteminfoBatchPutAPIRequest() *TaobaoUsceslIteminfoBatchPutAPIRequest {
	return poolTaobaoUsceslIteminfoBatchPutAPIRequest.Get().(*TaobaoUsceslIteminfoBatchPutAPIRequest)
}

// ReleaseTaobaoUsceslIteminfoBatchPutAPIRequest 将 TaobaoUsceslIteminfoBatchPutAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsceslIteminfoBatchPutAPIRequest(v *TaobaoUsceslIteminfoBatchPutAPIRequest) {
	v.Reset()
	poolTaobaoUsceslIteminfoBatchPutAPIRequest.Put(v)
}
