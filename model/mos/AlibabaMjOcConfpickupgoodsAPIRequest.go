package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcConfpickupgoodsAPIRequest 提货核销 API请求
// alibaba.mj.oc.confpickupgoods
//
// 此API用于在银泰商场中，消费者在提货中心提货时， 商户后台调用此接口进行提货核销操作
type AlibabaMjOcConfpickupgoodsAPIRequest struct {
	model.Params
	// 提货核销请求参数
	_confPickupGoodsRequest *ConfPickupGoodsReqDto
}

// NewAlibabaMjOcConfpickupgoodsRequest 初始化AlibabaMjOcConfpickupgoodsAPIRequest对象
func NewAlibabaMjOcConfpickupgoodsRequest() *AlibabaMjOcConfpickupgoodsAPIRequest {
	return &AlibabaMjOcConfpickupgoodsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjOcConfpickupgoodsAPIRequest) Reset() {
	r._confPickupGoodsRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcConfpickupgoodsAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.confpickupgoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcConfpickupgoodsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjOcConfpickupgoodsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfPickupGoodsRequest is ConfPickupGoodsRequest Setter
// 提货核销请求参数
func (r *AlibabaMjOcConfpickupgoodsAPIRequest) SetConfPickupGoodsRequest(_confPickupGoodsRequest *ConfPickupGoodsReqDto) error {
	r._confPickupGoodsRequest = _confPickupGoodsRequest
	r.Set("conf_pickup_goods_request", _confPickupGoodsRequest)
	return nil
}

// GetConfPickupGoodsRequest ConfPickupGoodsRequest Getter
func (r AlibabaMjOcConfpickupgoodsAPIRequest) GetConfPickupGoodsRequest() *ConfPickupGoodsReqDto {
	return r._confPickupGoodsRequest
}

var poolAlibabaMjOcConfpickupgoodsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjOcConfpickupgoodsRequest()
	},
}

// GetAlibabaMjOcConfpickupgoodsRequest 从 sync.Pool 获取 AlibabaMjOcConfpickupgoodsAPIRequest
func GetAlibabaMjOcConfpickupgoodsAPIRequest() *AlibabaMjOcConfpickupgoodsAPIRequest {
	return poolAlibabaMjOcConfpickupgoodsAPIRequest.Get().(*AlibabaMjOcConfpickupgoodsAPIRequest)
}

// ReleaseAlibabaMjOcConfpickupgoodsAPIRequest 将 AlibabaMjOcConfpickupgoodsAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjOcConfpickupgoodsAPIRequest(v *AlibabaMjOcConfpickupgoodsAPIRequest) {
	v.Reset()
	poolAlibabaMjOcConfpickupgoodsAPIRequest.Put(v)
}
