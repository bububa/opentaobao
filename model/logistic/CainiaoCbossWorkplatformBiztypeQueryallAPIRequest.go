package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCbossWorkplatformBiztypeQueryallAPIRequest
菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型 API请求
cainiao.cboss.workplatform.biztype.queryall

菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型。 目前调用者ISV */
type CainiaoCbossWorkplatformBiztypeQueryallAPIRequest struct {
	model.Params
	// level
	_level int64
	// tradeId
	_tradeId string
}

// New
