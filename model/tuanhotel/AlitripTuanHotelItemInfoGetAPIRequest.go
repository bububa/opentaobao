package tuanhotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelItemInfoGetAPIRequest 宝贝信息查询接口 API请求
// alitrip.tuan.hotel.item.info.get
//
// 商家查询发布的宝贝详情信息
type AlitripTuanHotelItemInfoGetAPIRequest struct {
	model.Params
	// 宝贝ID
	_itemId int64
}

// NewAlitripTuanHotelItemInfoGetRequest 初始化AlitripTuanHotelItemInfoGetAPIRequest对象
func NewAlitripTuanHotelItemInfoGetRequest() *AlitripTuanHotelItemInfoGetAPIRequest {
	return &AlitripTuanHotelItemInfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTuanHotelItemInfoGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelItemInfoGetAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.item.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelItemInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTuanHotelItemInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 宝贝ID
func (r *AlitripTuanHotelItemInfoGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripTuanHotelItemInfoGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolAlitripTuanHotelItemInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTuanHotelItemInfoGetRequest()
	},
}

// GetAlitripTuanHotelItemInfoGetRequest 从 sync.Pool 获取 AlitripTuanHotelItemInfoGetAPIRequest
func GetAlitripTuanHotelItemInfoGetAPIRequest() *AlitripTuanHotelItemInfoGetAPIRequest {
	return poolAlitripTuanHotelItemInfoGetAPIRequest.Get().(*AlitripTuanHotelItemInfoGetAPIRequest)
}

// ReleaseAlitripTuanHotelItemInfoGetAPIRequest 将 AlitripTuanHotelItemInfoGetAPIRequest 放入 sync.Pool
func ReleaseAlitripTuanHotelItemInfoGetAPIRequest(v *AlitripTuanHotelItemInfoGetAPIRequest) {
	v.Reset()
	poolAlitripTuanHotelItemInfoGetAPIRequest.Put(v)
}
