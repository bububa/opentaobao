package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeletevenueAPIRequest
大麦换验平台-第三方对外开放-场馆接口deleteVenue API请求
alibaba.damai.mev.open.deletevenue

开放接口，删除场馆 */
type AlibabaDamaiMevOpenDeletevenueAPIRequest struct {
	model.Params
	// 入参deleteVenueParam
	_deleteVenueParam *VenueIdOpenParam
}

// New
