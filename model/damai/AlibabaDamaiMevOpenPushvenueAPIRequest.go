package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushvenueAPIRequest
大麦换验平台-第三方对外开放-场馆接口pushVenue API请求
alibaba.damai.mev.open.pushvenue

开放接口推送场馆 */
type AlibabaDamaiMevOpenPushvenueAPIRequest struct {
	model.Params
	// 入参pushVenueParam
	_pushVenueParam *ThirdVenuePushOpenParam
}

// New
