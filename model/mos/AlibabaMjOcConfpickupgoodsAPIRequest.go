package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjocconfpickupgoodsAPIRequest 提货核销 API请求
// alibaba.mj.oc.confpickupgoods
//
// 此API用于在银泰商场中，消费者在提货中心提货时， 商户后台调用此接口进行提货核销操作
type AlibabamjocconfpickupgoodsAPIRequest struct {
	model.Params
	// 提货核销请求参数
	_confPickupGoodsRequest *ConfPickupGoodsReqDto
}

// NewAlibabamjocconfpickupgoodsRequest 初始化AlibabamjocconfpickupgoodsAPIRequest对象
func NewAlibabamjocconfpickupgoodsRequest() *AlibabamjocconfpickupgoodsAPIRequest {
	return &AlibabamjocconfpickupgoodsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjocconfpickupgoodsAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.confpickupgoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjocconfpickupgoodsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjocconfpickupgoodsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfPickupGoodsRequest is ConfPickupGoodsRequest Setter
// 提货核销请求参数
func (r *AlibabamjocconfpickupgoodsAPIRequest) SetConfPickupGoodsRequest(_confPickupGoodsRequest *ConfPickupGoodsReqDto) error {
	r._confPickupGoodsRequest = _confPickupGoodsRequest
	r.Set("conf_pickup_goods_request", _confPickupGoodsRequest)
	return nil
}

// GetConfPickupGoodsRequest ConfPickupGoodsRequest Getter
func (r AlibabamjocconfpickupgoodsAPIRequest) GetConfPickupGoodsRequest() *ConfPickupGoodsReqDto {
	return r._confPickupGoodsRequest
}
