package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxBiddingrefusesGetAPIRequest
tanx竞价失败反馈api API请求
taobao.tanx.biddingrefuses.get

竞价失败反馈根据创意id查询API提供 */
type TaobaoTanxBiddingrefusesGetAPIRequest struct {
	model.Params
	// dsp的创意id
	_creativeIds []string
	// dsp在tanx的memberid
	_memberId int64
	// dsp对应的tanx的token
	_token string
	// 1970年到现在的毫秒
	_signTime int64
	// 起始时间
	_startTime string
	// 截止时间
	_endTime string
}

// New
