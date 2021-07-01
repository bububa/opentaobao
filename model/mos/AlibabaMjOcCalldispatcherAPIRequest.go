package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcCalldispatcherAPIRequest
呼叫运力 API请求
alibaba.mj.oc.calldispatcher

定时达呼叫运力接口 */
type AlibabaMjOcCalldispatcherAPIRequest struct {
	model.Params
	// 入参
	_callDispatcherDTO *CallDispatcherDto
}

// New
