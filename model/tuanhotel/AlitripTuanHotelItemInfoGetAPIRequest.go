package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptuanhoteliteminfogetAPIRequest 宝贝信息查询接口 API请求
// alitrip.tuan.hotel.item.info.get
//
// 商家查询发布的宝贝详情信息
type AlitriptuanhoteliteminfogetAPIRequest struct {
	model.Params
	// 宝贝ID
	_itemId int64
}

// NewAlitriptuanhoteliteminfogetRequest 初始化AlitriptuanhoteliteminfogetAPIRequest对象
func NewAlitriptuanhoteliteminfogetRequest() *AlitriptuanhoteliteminfogetAPIRequest {
	return &AlitriptuanhoteliteminfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptuanhoteliteminfogetAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.item.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptuanhoteliteminfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptuanhoteliteminfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 宝贝ID
func (r *AlitriptuanhoteliteminfogetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitriptuanhoteliteminfogetAPIRequest) GetItemId() int64 {
	return r._itemId
}
