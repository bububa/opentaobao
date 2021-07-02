package mos

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcConfpickupgoodsAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.confpickupgoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcConfpickupgoodsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ConfPickupGoodsRequest Setter
// 提货核销请求参数
func (r *AlibabaMjOcConfpickupgoodsAPIRequest) SetConfPickupGoodsRequest(_confPickupGoodsRequest *ConfPickupGoodsReqDto) error {
	r._confPickupGoodsRequest = _confPickupGoodsRequest
	r.Set("conf_pickup_goods_request", _confPickupGoodsRequest)
	return nil
}

// Get ConfPickupGoodsRequest Getter
func (r AlibabaMjOcConfpickupgoodsAPIRequest) GetConfPickupGoodsRequest() *ConfPickupGoodsReqDto {
	return r._confPickupGoodsRequest
}
