package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallfuwuserviceitemlistAPIRequest 获取服务商品扩展信息 API请求
// tmall.fuwu.serviceitem.list
//
// 获取服务商品扩展信息
type TmallfuwuserviceitemlistAPIRequest struct {
	model.Params
	// 商品id列表，有数量限制
	_itemids []string
	// 商品所属卖家账号id
	_sellerId int64
}

// NewTmallfuwuserviceitemlistRequest 初始化TmallfuwuserviceitemlistAPIRequest对象
func NewTmallfuwuserviceitemlistRequest() *TmallfuwuserviceitemlistAPIRequest {
	return &TmallfuwuserviceitemlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallfuwuserviceitemlistAPIRequest) GetApiMethodName() string {
	return "tmall.fuwu.serviceitem.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallfuwuserviceitemlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallfuwuserviceitemlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemids is Itemids Setter
// 商品id列表，有数量限制
func (r *TmallfuwuserviceitemlistAPIRequest) SetItemids(_itemids []string) error {
	r._itemids = _itemids
	r.Set("itemids", _itemids)
	return nil
}

// GetItemids Itemids Getter
func (r TmallfuwuserviceitemlistAPIRequest) GetItemids() []string {
	return r._itemids
}

// SetSellerId is SellerId Setter
// 商品所属卖家账号id
func (r *TmallfuwuserviceitemlistAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TmallfuwuserviceitemlistAPIRequest) GetSellerId() int64 {
	return r._sellerId
}
