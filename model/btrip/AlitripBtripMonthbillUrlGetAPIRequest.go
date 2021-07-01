package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripMonthbillUrlGetAPIRequest
月账单数据查询 API请求
alitrip.btrip.monthbill.url.get

月账单数据查询 */
type AlitripBtripMonthbillUrlGetAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenAccountRq
}

// New
