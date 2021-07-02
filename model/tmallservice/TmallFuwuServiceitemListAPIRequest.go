package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuServiceitemListAPIRequest 获取服务商品扩展信息 API请求
// tmall.fuwu.serviceitem.list
//
// 获取服务商品扩展信息
type TmallFuwuServiceitemListAPIRequest struct {
	model.Params
	// 商品所属卖家账号id
	_sellerId int64
	// 商品id列表，有数量限制
	_itemids []int64
}

// NewTmallFuwuServiceitemListRequest 初始化TmallFuwuServiceitemListAPIRequest对象
func NewTmallFuwuServiceitemListRequest() *TmallFuwuServiceitemListAPIRequest {
	return &TmallFuwuServiceitemListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFuwuServiceitemListAPIRequest) GetApiMethodName() string {
	return "tmall.fuwu.serviceitem.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFuwuServiceitemListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSellerId is SellerId Setter
// 商品所属卖家账号id
func (r *TmallFuwuServiceitemListAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TmallFuwuServiceitemListAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// SetItemids is Itemids Setter
// 商品id列表，有数量限制
func (r *TmallFuwuServiceitemListAPIRequest) SetItemids(_itemids []int64) error {
	r._itemids = _itemids
	r.Set("itemids", _itemids)
	return nil
}

// GetItemids Itemids Getter
func (r TmallFuwuServiceitemListAPIRequest) GetItemids() []int64 {
	return r._itemids
}
