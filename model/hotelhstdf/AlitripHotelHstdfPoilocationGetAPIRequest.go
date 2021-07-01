package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfPoilocationGetAPIRequest
根据平台城市id分页查询poi location API请求
alitrip.hotel.hstdf.poilocation.get

根据平台城市id分页查询poi location */
type AlitripHotelHstdfPoilocationGetAPIRequest struct {
	model.Params
	// 参数封装
	_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam
}

// New
