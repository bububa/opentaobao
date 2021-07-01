package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTicketRuleQueryAPIRequest
【门票API2.0】门票规则信息查询接口 API请求
alitrip.ticket.rule.query

门票规则信息查询接口：返回商家上传的门票规则信息 */
type AlitripTicketRuleQueryAPIRequest struct {
	model.Params
	// 卖家景点规则编码
	_outRuleId string
}

// New
