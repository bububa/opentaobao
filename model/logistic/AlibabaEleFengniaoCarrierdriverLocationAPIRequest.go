package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoCarrierdriverLocationAPIRequest
查询骑手当前位置 API请求
alibaba.ele.fengniao.carrierdriver.location

查询骑手当前位置 */
type AlibabaEleFengniaoCarrierdriverLocationAPIRequest struct {
	model.Params
	// appid
	_appId string
	// 外部订单号
	_partnerOrderCode string
}

// New
