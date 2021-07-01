package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripUppcMemberGainAPIRequest
航司权益数据回流 API请求
alitrip.uppc.member.gain

航司权益数据回流 */
type AlitripUppcMemberGainAPIRequest struct {
	model.Params
	// 请求唯一标识
	_requestId string
	// 查询成功
	_statusCode int64
	// 权益截止时间（扩展字段）
	_responseJson string
	// 错误提示
	_errorMsg string
}

// New
