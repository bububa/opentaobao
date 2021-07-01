package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosPosAlarmAPIRequest
pos故障报警 API请求
alibaba.mos.pos.alarm

故障报警 */
type AlibabaMosPosAlarmAPIRequest struct {
	model.Params
	// 参数
	_param0 *PosLogDto
}

// New
