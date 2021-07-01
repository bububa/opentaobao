package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelTradeTemplateQueryAPIRequest
订单服务详情模版查询 API请求
alitrip.travel.trade.template.query

通过订单ID获取标注模版信息，商家可以根据模版来填充行业信息 */
type AlitripTravelTradeTemplateQueryAPIRequest struct {
	model.Params
	// 是否取最新的模版
	_isNew bool
	// 淘宝平台订单ID
	_orderId int64
}

// New
