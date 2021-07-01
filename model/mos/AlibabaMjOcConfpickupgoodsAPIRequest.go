package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcConfpickupgoodsAPIRequest
提货核销 API请求
alibaba.mj.oc.confpickupgoods

此API用于在银泰商场中，消费者在提货中心提货时， 商户后台调用此接口进行提货核销操作 */
type AlibabaMjOcConfpickupgoodsAPIRequest struct {
	model.Params
	// 提货核销请求参数
	_confPickupGoodsRequest *ConfPickupGoodsReqDto
}

// New
