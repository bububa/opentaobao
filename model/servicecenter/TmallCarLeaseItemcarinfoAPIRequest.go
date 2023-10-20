package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleaseitemcarinfoAPIRequest 整车租赁商品四级车型信息 API请求
// tmall.car.lease.itemcarinfo
//
// 整车租赁项目发布宝贝需要4级车型库，4级车型库信息需要回传
type TmallcarleaseitemcarinfoAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmallcarleaseitemcarinfoRequest 初始化TmallcarleaseitemcarinfoAPIRequest对象
func NewTmallcarleaseitemcarinfoRequest() *TmallcarleaseitemcarinfoAPIRequest {
	return &TmallcarleaseitemcarinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleaseitemcarinfoAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.itemcarinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleaseitemcarinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleaseitemcarinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallcarleaseitemcarinfoAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallcarleaseitemcarinfoAPIRequest) GetItemId() int64 {
	return r._itemId
}
