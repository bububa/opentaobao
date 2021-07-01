package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixSeatInfoQueryAPIRequest
分销商查询座位信息 API请求
alibaba.damai.maitix.seat.info.query

分销查询座位文案信息 */
type AlibabaDamaiMaitixSeatInfoQueryAPIRequest struct {
	model.Params
	// 入参
	_seatQueryParam *SeatQueryParam
}

// New
