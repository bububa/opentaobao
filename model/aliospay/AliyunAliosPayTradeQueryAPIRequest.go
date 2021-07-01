package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunAliosPayTradeQueryAPIRequest
查询支付结果接口 API请求
aliyun.alios.pay.trade.query

商户用来查询支付结果接口 */
type AliyunAliosPayTradeQueryAPIRequest struct {
	model.Params
	// 请求参数
	_queryTradeRequest *QueryTradeRequest
}

// New
