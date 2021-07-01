package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrSoldOrderlistQueryJstAPIRequest
已入塔商家查询订单列表 API请求
tmall.nr.sold.orderlist.query.jst

该服务用于已入聚石塔的商家，获取订单列表信息； */
type TmallNrSoldOrderlistQueryJstAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *NrTimingOrderSoldQueryReqDto
}

// New
