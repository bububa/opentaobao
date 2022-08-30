package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeOrderdetailEntryGetAPIRequest 获取订单详情页跳转入口配置 API请求
// taobao.opentrade.orderdetail.entry.get
//
// 获取订单详情页跳转入口配置
type TaobaoOpentradeOrderdetailEntryGetAPIRequest struct {
	model.Params
	// 商品id列表
	_itemIdList int64
}

// NewTaobaoOpentradeOrderdetailEntryGetRequest 初始化TaobaoOpentradeOrderdetailEntryGetAPIRequest对象
func NewTaobaoOpentradeOrderdetailEntryGetRequest() *TaobaoOpentradeOrderdetailEntryGetAPIRequest {
	return &TaobaoOpentradeOrderdetailEntryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeOrderdetailEntryGetAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.orderdetail.entry.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeOrderdetailEntryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemIdList is ItemIdList Setter
// 商品id列表
func (r *TaobaoOpentradeOrderdetailEntryGetAPIRequest) SetItemIdList(_itemIdList int64) error {
	r._itemIdList = _itemIdList
	r.Set("item_id_list", _itemIdList)
	return nil
}

// GetItemIdList ItemIdList Getter
func (r TaobaoOpentradeOrderdetailEntryGetAPIRequest) GetItemIdList() int64 {
	return r._itemIdList
}
