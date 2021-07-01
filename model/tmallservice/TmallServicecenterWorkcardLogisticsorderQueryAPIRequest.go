package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardLogisticsorderQueryAPIRequest
物流单信息查询 API请求
tmall.servicecenter.workcard.logisticsorder.query

物流订单信息查询API */
type TmallServicecenterWorkcardLogisticsorderQueryAPIRequest struct {
	model.Params
	// 物流单号
	_logisticsOrderId int64
}

// New
