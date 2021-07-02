package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemGetAPIRequest 根据id查询商品 API请求
// taobao.scitem.get
//
// 根据id查询商品
type TaobaoScitemGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTaobaoScitemGetRequest 初始化TaobaoScitemGetAPIRequest对象
func NewTaobaoScitemGetRequest() *TaobaoScitemGetAPIRequest {
	return &TaobaoScitemGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoScitemGetAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoScitemGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoScitemGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoScitemGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
