package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmMemberScoreGetAPIRequest
获取会员某时间段积分 API请求
alibaba.westcrm.member.score.get

获取会员某时间段积分 */
type AlibabaWestcrmMemberScoreGetAPIRequest struct {
	model.Params
	// requestId
	_requestId string
	// 支付宝id
	_alipayId string
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
}

// New
